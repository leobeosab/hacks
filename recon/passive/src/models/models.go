package models

type Scan struct {
	Name            string   `json:"name"`
	Subdomains      []Target `json:"subdomains"`
	DNSWordlistPath string   `json:"dnsWordlistPath"`
}

type Target struct {
	Root     string   `json:"root"`
	Wildcard bool     `json:"wildcard"`
	Domains  []Domain `json:"domains"`
}

type Domain struct {
	Name string `json:"name"`
}
