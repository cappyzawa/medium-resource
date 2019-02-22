package resource

import "github.com/Medium/medium-sdk-go"

// MediumClient can access to medium
//go:generate counterfeiter -o fakes/medium_client.go MediumClient
type MediumClient interface {
	// https://godoc.org/github.com/Medium/medium-sdk-go#Medium.CreatePost
	CreatePost(o medium.CreatePostOptions) (*medium.Post, error)
	// https://godoc.org/github.com/Medium/medium-sdk-go#Medium.GetUser
	GetUser(userID string) (*medium.User, error)
}
