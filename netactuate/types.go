package netactuate

type ZoneSummary struct {
	Name string `json:"name"`
	Type string `json:"type"`
	ID   int    `json:"id"`
	TTL  int    `json:"ttl"`
}

type ZoneList struct {
	Result  string        `json:"result"`
	Message string        `json:"message"`
	Data    []ZoneSummary `json:"data"`
	Code    int           `json:"code"`
}

type DNSRecordPostResponseData struct {
	ZoneType string `json:"type"`
	Name     string `json:"name"`
	Content  string `json:"content"`
	DomainID int    `json:"domain_id"`
	TTL      int    `json:"ttl"`
	ID       int    `json:"id"`
}

type DNSRecordPostResponse struct {
	Result string                    `json:"result"`
	Data   DNSRecordPostResponseData `json:"data"`
	Code   int                       `json:"code"`
}

type DNSRecord struct {
	Name       string `json:"name"`
	RecordType string `json:"type"`
	Content    string `json:"content"`
	ID         int    `json:"id"`
	TTL        int    `json:"ttl"`
}

type DNSRecordListResponse struct {
	Result  string      `json:"result"`
	Message string      `json:"message"`
	Data    []DNSRecord `json:"data"`
	Code    int         `json:"code"`
}
