package adapter

import (
	"net"
	"reflect"
	"testing"

	"github.com/dnguy078/go-detector/models"
)

const sourceFile = "../data/city_seed_data.mmdb"

func TestNewGeoIP(t *testing.T) {
	tests := []struct {
		name     string
		filePath string
		wantErr  bool
	}{
		{
			name:     "success",
			filePath: sourceFile,
		},
		{
			name:     "invalid path",
			filePath: "../data/invalid",
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewGeoIP(tt.filePath)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewGeoIP() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGeoIP_Location(t *testing.T) {
	tests := []struct {
		name             string
		wantErr          bool
		ipAddress        string
		expectedLocation *models.Location
	}{
		{
			name:      "success",
			ipAddress: "206.81.252.6",
			expectedLocation: &models.Location{
				Lat:    39.2293,
				Lon:    -76.6907,
				Radius: 20,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g, err := NewGeoIP(sourceFile)
			if err != nil {
				t.Fatal(err)
			}

			ip := net.ParseIP(tt.ipAddress)

			loc, err := g.Location(ip)
			if (err != nil) != tt.wantErr {
				t.Errorf("TestGeoIP_Location() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !reflect.DeepEqual(loc, tt.expectedLocation) {
				t.Errorf("TestGeoIP_Location() got location = %+v, want %+v", loc, tt.expectedLocation)
			}
		})
	}
}
