package resource

// Source is source configuration for resource
type Source struct {
	AccessToken string `json:"access_token"`
}

// Version is primary key for resource
type Version struct {
	ID string `json:"id"`
}

// MetadataPair is key-value for putting resource
type MetadataPair struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}
