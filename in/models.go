package in

import (
	"github.com/cappyzawa/medium-resource"
)

type Request struct {
	Source  resource.Source
	Params  Params
	Version resource.Version
}

type Params struct {
}

type Response struct {
	Version  resource.Version
	Metadata []resource.MetadataPair
}
