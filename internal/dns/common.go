package dns

type DNS_Record struct {
	Comment  string   `json:"comment"`
	Type     string   `json:"type"`
	Content  string   `json:"content"`
	Name     string   `json:"name"`
	Priority int      `json:"priority"`
	Proxied  bool     `json:"proxied"`
	Tags     []string `json:"tags"`
	TTL      int      `json:"ttl"`
	Id       string   `json:"id,omitempty"`
}
