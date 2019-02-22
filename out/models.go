package out

import "github.com/cappyzawa/medium-resource"

type Request struct {
	Source resource.Source `json:"source"`
	Params Params          `json:"params"`
}

type Params struct {
	Format       string   `json:"format"`
	File         string   `json:"file"`
	Tags         []string `json:"tags"`
	CanonicalURL string   `json:"canonical_url"`
	Status       string   `json:"status"`
	License      string   `json:"license"`
}

type Response struct {
	Version  resource.Version
	Metadata []resource.MetadataPair
}
