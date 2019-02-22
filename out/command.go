package out

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/Medium/medium-sdk-go"
	"github.com/cappyzawa/medium-resource"
)

type Command struct {
	MediumClient resource.MediumClient
}

// Run publish an article based on request.
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
		ContentFormat: medium.ContentFormatMarkdown,
		Tags:          []string{},
		CanonicalURL:  "",
		PublishStatus: medium.PublishStatusDraft,
		License:       "",
	}
	title, content, err := c.ExtractTitleAndContent(fmt.Sprintf("%s/%s", sourceDir, request.Params.ContentFile))
	if err != nil {
		return nil, err
	}
	o.Content = content
	o.Title = title
	if request.Params.Title != "" {
		o.Title = request.Params.Title
	}
	if request.Params.Format != "" {
		o.ContentFormat = medium.ContentFormat(request.Params.Format)
	}
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

func (c *Command) ExtractTitleAndContent(path string) (string, string, error) {
	contents, err := ioutil.ReadFile(path)
	if err != nil {
		return "", "", err
	}
	separated := strings.Split(string(contents), "\n")
	title := strings.TrimLeft(separated[0], "# ")
	content := strings.Join(separated[1:], "\n")
	return title, content, err
}
