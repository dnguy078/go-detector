package service

import (
	"net"
	"reflect"
	"testing"
	"time"

	"github.com/dnguy078/go-detector/models"
)

type fakeStorage struct {
	insertErr error
	prevResp  *models.UserGeoEvent
	prevErr   error

	nextResp *models.UserGeoEvent
	nextErr  error
}

func (fs *fakeStorage) InsertUserEvent(*models.UserGeoEvent) error {
	return fs.insertErr
}
func (fs *fakeStorage) GetPreviousIPAccess(*models.UserGeoEvent) (*models.UserGeoEvent, error) {
	return fs.prevResp, fs.prevErr
}
func (fs *fakeStorage) GetSubsequentIPAccess(*models.UserGeoEvent) (*models.UserGeoEvent, error) {
	return fs.nextResp, fs.nextErr
}
func (fs *fakeStorage) Close() error {
	return nil
}

type fakeGeo struct {
	locResp *models.Location
	locErr  error
}

func (fg *fakeGeo) Location(net.IP) (*models.Location, error) {
	return fg.locResp, fg.locErr
}
func (fg *fakeGeo) Close() error {
	return nil
}

func Test_distance(t *testing.T) {
	type args struct {
		a *models.Location
		b *models.Location
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "success",
			args: args{
				a: &models.Location{
					Lat:    39.1702,
					Lon:    76,
					Radius: 20,
				},
				b: &models.Location{
					Lat:    30,
					Lon:    -97,
					Radius: 5,
				},
			},
			want: 7634.965413899729,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := distance(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("distance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSpeed(t *testing.T) {
	type args struct {
		a          *models.Location
		b          *models.Location
		timestampA time.Time
		timestampB time.Time
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "preceding",
			args: args{
				a: &models.Location{
					Lat:    39.1702,
					Lon:    -76.8538,
					Radius: 20,
				},
				b: &models.Location{
					Lat:    30.3764,
					Lon:    -97.7078,
					Radius: 5,
				},
				timestampA: time.Unix(1514764800, 0),
				timestampB: time.Unix(1514764800, 0),
			},
			want: 1325701,
		},
		{
			name: "subsequent",
			args: args{
				a: &models.Location{
					Lat:    39.1702,
					Lon:    -76.8538,
					Radius: 20,
				},
				b: &models.Location{
					Lat:    34.0494,
					Lon:    -118.2641,
					Radius: 200,
				},
				timestampA: time.Unix(1514764800, 0),
				timestampB: time.Unix(1514851200, 0),
			},
			want: 95,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := speed(tt.args.a, tt.args.b, tt.args.timestampA, tt.args.timestampB); got != tt.want {
				t.Errorf("Speed() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_geoSuspiciousSvc_Suspicious(t *testing.T) {
	type fields struct {
		db  LoginStorage
		geo GeoIPer
	}
	type args struct {
		req models.Request
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *models.Response
		wantErr bool
	}{
		{
			name: "success",
			fields: fields{
				db: &fakeStorage{
					prevResp: &models.UserGeoEvent{
						Username:   "test",
						EventUUID:  "test",
						InsertedAt: 1574733772,
						Timestamp:  1574733772,
						IPAddress:  "test",
						Location: &models.Location{
							Lat:    33.7854,
							Lon:    -117.7948,
							Radius: 5,
						},
					},
				},
				geo: &fakeGeo{
					locResp: &models.Location{
						Lat:    41.8797,
						Lon:    -87.6435,
						Radius: 200,
					},
				},
			},
			args: args{
				req: models.Request{
					Username:      "integration-test",
					UnixTimestamp: 1574733772, // 11/26/2019 @ 2:02am (UTC))
					EventUUID:     "request",
					IPAddress:     "37.60.254.143", // Chicago, IL
				},
			},
			want: &models.Response{
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
					Speed:     1726950,
					IP:        "test",
					Timestamp: 1574733772,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gs := &geoSuspiciousSvc{
				db:  tt.fields.db,
				geo: tt.fields.geo,
			}
			got, err := gs.Suspicious(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("geoSuspiciousSvc.Suspicious() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("geoSuspiciousSvc.Suspicious() = %v, want %v", got, tt.want)
				t.Errorf("PreviousIP Address got \n%+v\nexpected: \n%+v\n", got.PrecedingIPAccess, tt.want.PrecedingIPAccess)
			}
		})
	}
}
