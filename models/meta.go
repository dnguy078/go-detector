package models

// Loginmeta holds login metadata
type Loginmeta struct {
	Username      string `json:"username"`
	UnixTimestamp int    `json:"unix_timestamp"`
	EventUUID     string `json:"event_uuid"`
	IPAddress     string `json:"ip_address"`
}
