package models

import "time"

// UserGeoEvent holds a login request metadata and Location corresponding to the event
type UserGeoEvent struct {
	Username   string
	EventUUID  string
	InsertedAt int64
	Timestamp  int64
	IPAddress  string
	*Location
}

// NewUserGeoEvent returns a UserGeoEvent
func NewUserGeoEvent(req Request, loc *Location) *UserGeoEvent {
	return &UserGeoEvent{
		Username:   req.Username,
		EventUUID:  req.EventUUID,
		InsertedAt: time.Now().Unix(),
		Timestamp:  req.UnixTimestamp,
		IPAddress:  req.IPAddress,
		Location:   loc,
	}
}
