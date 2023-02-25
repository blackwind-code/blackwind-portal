package dns

type DNS_Record_Create struct {
	*DNS_Record
}

type DNS_Record_Update struct {
	Comment string `json: "comment,omitempty"`
	Content string `json: "content,omitempty"`
}


