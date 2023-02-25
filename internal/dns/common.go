package dns

type DNS_Record struct {
	Comment  string   `json:"comment,omitempty"`
	Type     string   `json:"type,omitempty"`
	Content  string   `json:"content,omitempty"`
	Name     string   `json:"name,omitempty"`
	Priority int      `json:"priority,omitempty"`
	Proxied  bool     `json:"proxied,omitempty"`
	Tags     []string `json:"tags,omitempty"`
	TTL      int      `json:"ttl,omitempty"`
	Id       string   `json:"omitempty"`
}
