package out

import "github.com/cappyzawa/medium-resource"

// Request is payload input to out.
type Request struct {
	Source resource.Source `json:"source"`
	Params Params          `json:"params"`
}

// Params is parameters for put.
type Params struct {
	ContentFile  string   `json:"content_file"`
	Tags         []string `json:"tags"`
	CanonicalURL string   `json:"canonical_url"`
	Status       string   `json:"status"`
	License      string   `json:"license"`
}

// Response outputs version and metadata for resource.
type Response struct {
	Version  resource.Version
	Metadata []resource.MetadataPair
}
