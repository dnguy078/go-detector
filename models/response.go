package models

// Response holds current location, suspicious travels, previous and subsequent ip access
type Response struct {
	CurrentGeo                     *Location `json:"currentGeo"`
	TravelToCurrentGeoSuspicious   bool      `json:"travelToCurrentGeoSuspicious"`
	TravelFromCurrentGeoSuspicious bool      `json:"travelFromCurrentGeoSuspicious"`
	PrecedingIPAccess              *IPAccess `json:"precedingIpAccess,omitempty"`
	SubsequentIPAccess             *IPAccess `json:"subsequentIpAccess,omitempty"`
}
