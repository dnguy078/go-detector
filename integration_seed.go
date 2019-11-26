// +build integration

package main

import (
	"github.com/dnguy078/go-detector/models"
	"github.com/google/uuid"
)

var requests = []models.Request{
	models.Request{
		Username:      "integration-test",
		UnixTimestamp: 1574733413, // 11/26/2019 @ 1:56am (UTC)
		EventUUID:     uuid.New().String(),
		IPAddress:     "172.90.83.116", // Orange, CA
	},
	models.Request{
		Username:      "integration-test",
		UnixTimestamp: 1574733772, // 11/26/2019 @ 2:02am (UTC))
		EventUUID:     uuid.New().String(),
		IPAddress:     "37.60.254.143", // Chicago, IL
	},
	models.Request{
		Username:      "integration-test",
		UnixTimestamp: 1574735184, // 11/26/2019 @ 2:26am (UTC)
		EventUUID:     uuid.New().String(),
		IPAddress:     "107.77.66.81", // Houston, Texas
	},
	models.Request{
		Username:      "integration-test",
		UnixTimestamp: 1574733420, // 11/26/2019 @ 1:57am (UTC))
		EventUUID:     uuid.New().String(),
		IPAddress:     "204.89.92.153", // Tallahassee, Florida
	},
}

var responses = []models.Response{
	// Orange, CA
	models.Response{
		CurrentGeo: &models.Location{
			Lat:    33.7854,
			Lon:    -117.7948,
			Radius: 5,
		},
		TravelToCurrentGeoSuspicious:   false,
		TravelFromCurrentGeoSuspicious: false,
	},
	// Orange,CA to Chicago,IL
	models.Response{
		CurrentGeo: &models.Location{
			Lat:    41.8797,
			Lon:    -87.6435,
			Radius: 200,
		},
		TravelToCurrentGeoSuspicious:   true,
		TravelFromCurrentGeoSuspicious: false,
		PrecedingIPAccess: &models.IPAccess{
			Location: models.Location{
				Lat:    33.7854,
				Lon:    -117.7948,
				Radius: 5,
			},
			Speed:     17317,
			IP:        "172.90.83.116",
			Timestamp: 1574733413,
		},
	},
	// Chicago,IL to Houston,TX
	models.Response{
		CurrentGeo: &models.Location{
			Lat:    29.772,
			Lon:    -95.3644,
			Radius: 100,
		},
		TravelToCurrentGeoSuspicious:   true,
		TravelFromCurrentGeoSuspicious: false,
		PrecedingIPAccess: &models.IPAccess{
			Location: models.Location{
				Lat:    41.8797,
				Lon:    -87.6435,
				Radius: 200,
			},
			Speed:     2398,            // miles per hour
			IP:        "37.60.254.143", // Chicago, IL
			Timestamp: 1574733772,
		},
	},
	// Orange -> Tallahassee -> Chicago
	models.Response{
		CurrentGeo: &models.Location{
			Lat:    30.4369,
			Lon:    -84.2763,
			Radius: 1000,
		},
		TravelToCurrentGeoSuspicious:   true,
		TravelFromCurrentGeoSuspicious: true,
		// Orange
		PrecedingIPAccess: &models.IPAccess{
			Location: models.Location{
				Lat:    33.7854,
				Lon:    -117.7948,
				Radius: 5,
			},
			Speed:     1011182,
			IP:        "172.90.83.116",
			Timestamp: 1574733413,
		},
		// Chicago
		SubsequentIPAccess: &models.IPAccess{
			Location: models.Location{
				Lat:    41.8797,
				Lon:    -87.6435,
				Radius: 200,
			},
			IP:        "37.60.254.143",
			Timestamp: 1574733772,
			Speed:     8307,
		},
	},
}
