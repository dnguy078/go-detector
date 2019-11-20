package models

type Response struct {
	CurrentGeo struct {
		Lat    float64 `json:"lat"`
		Lon    float64 `json:"lon"`
		Radius int     `json:"radius"`
	} `json:"currentGeo"`
	TravelToCurrentGeoSuspicious   bool `json:"travelToCurrentGeoSuspicious"`
	TravelFromCurrentGeoSuspicious bool `json:"travelFromCurrentGeoSuspicious"`
	PrecedingIPAccess              struct {
		IP        string  `json:"ip"`
		Speed     int     `json:"speed"`
		Lat       float64 `json:"lat"`
		Lon       float64 `json:"lon"`
		Radius    int     `json:"radius"`
		Timestamp int     `json:"timestamp"`
	} `json:"precedingIpAccess"`
	SubsequentIPAccess struct {
		IP        string  `json:"ip"`
		Speed     int     `json:"speed"`
		Lat       float64 `json:"lat"`
		Lon       float64 `json:"lon"`
		Radius    int     `json:"radius"`
		Timestamp int     `json:"timestamp"`
	} `json:"subsequentIpAccess"`
}
