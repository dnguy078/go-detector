package adapter

import (
	"net"

	"github.com/dnguy078/go-detector/models"
	geo "github.com/oschwald/geoip2-golang"
)

// GeoIP is a wrapper around the MaxMind GeoIP2 and GeoLite2 databases, loads from seed file and runs inmem
type GeoIP struct {
	reader *geo.Reader
}

// NewGeoIP returns a GeoIP object
func NewGeoIP(seedFilePath string) (*GeoIP, error) {
	r, err := geo.Open(seedFilePath)
	if err != nil {
		return nil, err
	}

	return &GeoIP{
		reader: r,
	}, nil
}

func (g *GeoIP) Close() error {
	return g.reader.Close()
}

// Location returns the location based upon ip address
func (g *GeoIP) Location(ipAddress net.IP) (*models.Location, error) {
	city, err := g.reader.City(ipAddress)
	if err != nil {
		return nil, err
	}
	return &models.Location{
		Lat:    city.Location.Latitude,
		Lon:    city.Location.Longitude,
		Radius: int(city.Location.AccuracyRadius),
	}, nil
}
