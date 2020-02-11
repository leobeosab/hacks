package models

type Scan struct {
	Name            string            `json:"name"`
	Subdomains      map[string]Target `json:"subdomains"`
	DNSWordlistPath string            `json:"dnsWordlistPath"`
}

type Target struct {
	Root     string            `json:"root"`
	Wildcard bool              `json:"wildcard"`
	Domains  map[string]Domain `json:"domains"`
}

type Domain struct {
	Name string `json:"name"`
}
