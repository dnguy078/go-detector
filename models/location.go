package models

// Location holds location information
type Location struct {
	Lat    float64 `json:"lat"`
	Lon    float64 `json:"lon"`
	Radius int     `json:"radius"`
}

// IPAccess holds location along with speeds from previous/subsequent locations, and timestamp
type IPAccess struct {
	Location
	IP        string `json:"ip"`
	Speed     int    `json:"speed"`
	Timestamp int64  `json:"timestamp"`
}
