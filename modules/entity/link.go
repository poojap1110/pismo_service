package entity

// Link hateoas links added to responses
type Link struct {
	Rel  string `json:"rel"`
	HREF string `json:"href"`
}

// LinkHost declares the host in the HREF field when calling NewLink()
var LinkHost string

// NewLink creates a new link based on the LinkHost
func NewLink(url string, rel string) Link {
	return Link{rel, LinkHost + url}
}
