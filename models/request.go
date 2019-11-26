package models

// Request holds login metadata
type Request struct {
	Username      string `json:"username"`
	UnixTimestamp int64  `json:"unix_timestamp"`
	EventUUID     string `json:"event_uuid"`
	IPAddress     string `json:"ip_address"`
}

// Validate validates incoming request
func (r Request) Validate() error {
	panic("Implement me!")
}
