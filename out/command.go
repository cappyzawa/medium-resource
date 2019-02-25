package out

import (
	"errors"
	"fmt"
	"io/ioutil"
	"path"
	"regexp"
	"strings"

	"github.com/Medium/medium-sdk-go"
	"github.com/cappyzawa/medium-resource"
)

var r *regexp.Regexp

func init() {
	r = regexp.MustCompile(`<title>.*?</title>`)
}

// Command has MediumClient for posting a content.
type Command struct {
	MediumClient resource.MediumClient
}

// Run publishes an article based on request.
func (c *Command) Run(sourceDir string, request Request) (*Response, error) {
	if request.Source.AccessToken == "" {
		return nil, errors.New("\"access_token\" is missing")
	}
	if request.Params.ContentFile == "" {
		return nil, errors.New("\"content_file\" is missing")
	}
	u, err := c.MediumClient.GetUser("")
	if err != nil {
		return nil, err
	}

	// https://godoc.org/github.com/Medium/medium-sdk-go#CreatePostOptions
	o := medium.CreatePostOptions{
		UserID:        u.ID,
		Title:         "",
		Content:       "",
		ContentFormat: "",
		Tags:          []string{},
		CanonicalURL:  "",
		PublishStatus: medium.PublishStatusDraft,
		License:       "",
	}
	format, title, content, err := c.ExtractFromFile(fmt.Sprintf("%s/%s", sourceDir, request.Params.ContentFile))
	if err != nil {
		return nil, err
	}
	o.Content = content
	o.Title = title
	o.ContentFormat = format
	if len(request.Params.Tags) != 0 {
		o.Tags = append(o.Tags, request.Params.Tags...)
	}
	if request.Params.CanonicalURL != "" {
		o.CanonicalURL = request.Params.CanonicalURL
	}
	if request.Params.Status != "" {
		o.PublishStatus = medium.PublishStatus(request.Params.Status)
	}
	if request.Params.License != "" {
		o.License = medium.License(request.Params.License)
	}

	posted, err := c.MediumClient.CreatePost(o)
	if err != nil {
		return nil, err
	}

	return &Response{
		resource.Version{ID: posted.ID},
		[]resource.MetadataPair{
			{"title", posted.Title},
			{"status", string(posted.PublishState)},
			{"url", posted.URL},
		},
	}, nil
}

// ExtractFromFile extracts format, title, content from file.
func (c *Command) ExtractFromFile(filePath string) (medium.ContentFormat, string, string, error) {
	contents, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", "", "", err
	}
	switch path.Ext(filePath) {
	case ".md":
		title, content, err := c.extractTitleAndContentByMd(contents)
		if err != nil {
			return "", "", "", err
		}
		return medium.ContentFormat(medium.ContentFormatMarkdown), title, content, nil
	case ".html":
		title, content, err := c.extractTitleAndContentByHtml(contents)
		if err != nil {
			return "", "", "", err
		}
		return medium.ContentFormatHTML, title, content, nil
	default:
		return "", "", "", errors.New("no support ext: " + path.Ext(filePath))
	}
}

func (c *Command) extractTitleAndContentByMd(contents []byte) (string, string, error) {
	separated := strings.Split(string(contents), "\n")
	title := strings.TrimLeft(separated[0], "# ")
	content := strings.Join(separated[1:], "\n")
	return title, content, nil
}

func (c *Command) extractTitleAndContentByHtml(contents []byte) (string, string, error) {
	matched := r.FindSubmatch(contents)
	// matched = <title>title</title>
	title := strings.TrimSuffix(strings.TrimPrefix(string(matched[0]), "<title>"), "</title>")
	if title == "" {
		return "", "", errors.New("title tag is required")
	}
	return title, string(contents), nil
}
