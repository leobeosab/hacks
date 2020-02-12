package models

type Scan struct {
	Name                string            `json:"name"`
	Subdomains          map[string]Target `json:"subdomains"`
	DNSWordlistPath     string            `json:"dnsWordlistPath"`
	DirbustWordlistPath string            `json:"dirbustWordlistPath"`
}

type Target struct {
	Root     string            `json:"root"`
	Wildcard bool              `json:"wildcard"`
	Domains  map[string]Domain `json:"domains"`
}

type Domain struct {
	Name        string          `json:"name"`
	DirbResults []DirBustResult `json:"dirbResults"`
}

type DirBustResult struct {
	Path   string `json:"path"`
	Status int    `json:"status"`
}
