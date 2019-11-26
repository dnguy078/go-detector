package service

import (
	"math"
	"net"
	"time"

	"github.com/dnguy078/go-detector/models"

	"github.com/umahmood/haversine"
)

const (
	// SUPERMAN_SPEED is the speed threshold for determining of a login is supicious
	SUPERMAN_SPEED = 500
)

// geoSuspiciousSvc determines of a login attempt is suspicious
type geoSuspiciousSvc struct {
	db  LoginStorage
	geo GeoIPer
}

// NewGeoSuspiciousSvc returns a geoSuspiciousSvc
func NewGeoSuspiciousSvc(db LoginStorage, geoIP GeoIPer) *geoSuspiciousSvc {
	return &geoSuspiciousSvc{
		db:  db,
		geo: geoIP,
	}
}

// LoginStorage read/writes locations
type LoginStorage interface {
	InsertUserEvent(*models.UserGeoEvent) error
	GetPreviousIPAccess(*models.UserGeoEvent) (*models.UserGeoEvent, error)
	GetSubsequentIPAccess(*models.UserGeoEvent) (*models.UserGeoEvent, error)
	Close() error
}

// GeoIPer returns locations based upon IP
type GeoIPer interface {
	Location(net.IP) (*models.Location, error)
	Close() error
}

// Suspicious determines if login attempts are suspicious
func (gs *geoSuspiciousSvc) Suspicious(req models.Request) (*models.Response, error) {
	ip := net.ParseIP(req.IPAddress)
	curr, err := gs.geo.Location(ip)
	if err != nil {
		return nil, err
	}

	e := models.NewUserGeoEvent(req, curr)
	prev, next, err := gs.prevAndNextEvent(e)
	if err != nil {
		return nil, err
	}
	if err := gs.db.InsertUserEvent(e); err != nil {
		return nil, err
	}

	return gs.buildResponse(e, prev, next)
}

func (gs *geoSuspiciousSvc) prevAndNextEvent(e *models.UserGeoEvent) (*models.UserGeoEvent, *models.UserGeoEvent, error) {
	prev, next := new(models.UserGeoEvent), new(models.UserGeoEvent)
	prev, err := gs.db.GetPreviousIPAccess(e)
	if err != nil {
		return nil, nil, err
	}
	next, err = gs.db.GetSubsequentIPAccess(e)

	return prev, next, err
}

func (gs *geoSuspiciousSvc) buildResponse(curr, prev, next *models.UserGeoEvent) (*models.Response, error) {
	resp := &models.Response{
		CurrentGeo: curr.Location,
	}

	if prev != nil {
		speed := speed(curr.Location, prev.Location, time.Unix(curr.Timestamp, 0), time.Unix(prev.Timestamp, 0))
		resp.PrecedingIPAccess = &models.IPAccess{
			Location: models.Location{
				Lat:    prev.Lat,
				Lon:    prev.Lon,
				Radius: prev.Radius,
			},
			Speed:     speed,
			Timestamp: prev.Timestamp,
			IP:        prev.IPAddress,
		}
		resp.TravelToCurrentGeoSuspicious = supermanSpeed(speed)
	}

	if next != nil {
		speed := speed(curr.Location, next.Location, time.Unix(curr.Timestamp, 0), time.Unix(next.Timestamp, 0))
		resp.SubsequentIPAccess = &models.IPAccess{
			Location: models.Location{
				Lat:    next.Lat,
				Lon:    next.Lon,
				Radius: next.Radius,
			},
			Speed:     speed,
			Timestamp: next.Timestamp,
			IP:        next.IPAddress,
		}
		resp.TravelFromCurrentGeoSuspicious = supermanSpeed(speed)
	}

	return resp, nil
}

func supermanSpeed(speed int) bool {
	return speed > SUPERMAN_SPEED
}

// distance returns the distance between two locations using haversine formula
func distance(a, b *models.Location) float64 {
	locA := haversine.Coord{Lat: a.Lat, Lon: a.Lon}
	locB := haversine.Coord{Lat: b.Lat, Lon: b.Lon}
	m, _ := haversine.Distance(locA, locB)
	
	return m
}

// Speed returns the miles per hours traveled between two distance and time
func speed(a, b *models.Location, timestampA, timestampB time.Time) int {
	// kmToMiles := float64(a.Radius+b.Radius) * .64, unsure how radius is used to determine speed...
	d := distance(a, b)
	timeDiff := math.Abs(timestampB.Sub(timestampA).Hours())

	if timeDiff == 0 {
		timeDiff = .001
	}

	return int(d / timeDiff)
}
