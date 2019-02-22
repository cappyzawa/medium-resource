package out

import (
	"github.com/Medium/medium-sdk-go"
	"github.com/cappyzawa/medium-resource"
)

type Command struct {
	MediumClient resource.MediumClient
}

// Run publish an article based on request.
func (c *Command) Run(sourceDir string, request Request) (*Response, error) {
	u, err := c.MediumClient.GetUser("")
	if err != nil {
		return nil, err
	}
	// https://godoc.org/github.com/Medium/medium-sdk-go#CreatePostOptions
	o := medium.CreatePostOptions{
		UserID:  u.ID,
		Title:   "title",
		Content: "#hoge",
		// https://godoc.org/github.com/Medium/medium-sdk-go#ContentFormat
		ContentFormat: medium.ContentFormat(request.Params.Format),
		Tags:          request.Params.Tags,
		CanonicalURL:  request.Params.CanonicalURL,
		// https://godoc.org/github.com/Medium/medium-sdk-go#PublishStatus
		PublishStatus: medium.PublishStatus(request.Params.Status),
		// https://godoc.org/github.com/Medium/medium-sdk-go#License
		License: medium.License(request.Params.License),
	}
	_, err = c.MediumClient.CreatePost(o)
	if err != nil {
		return nil, err
	}

	return &Response{
		resource.Version{},
		[]resource.MetadataPair{},
	}, nil
}
