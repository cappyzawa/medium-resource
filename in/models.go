package in

import (
	"github.com/cappyzawa/medium-resource"
)

// Request is payload input to in.
type Request struct {
	Source  resource.Source
	Params  Params
	Version resource.Version
}

// Params is parameters for get.
type Params struct {
}

// Response outputs version and metadata for resource.
type Response struct {
	Version  resource.Version
	Metadata []resource.MetadataPair
}
