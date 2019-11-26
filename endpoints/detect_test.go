package endpoints

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/dnguy078/go-detector/models"
)

type fakeDetection struct {
	resp *models.Response
	err  error
}

func (f *fakeDetection) Suspicious(models.Request) (*models.Response, error) {
	return f.resp, f.err
}

func TestDetect(t *testing.T) {
	cases := []struct {
		name    string
		payload string
		err     error
		svc     DetectionService

		expectedStatusCode int
		expectedResponse   string
	}{
		{
			name: "success",
			payload: `{
				"username":"bob",
				"unix_timestamp":1514764800,
				"event_uuid":"85ad929a-db03-4bf4-9541-8f728fa12e42",
				"ip_address":"206.81.252.6"
			 }`,
			svc: &fakeDetection{
				resp: &models.Response{
					CurrentGeo: &models.Location{
						Lat: 1,
						Lon: 2,
					},
					TravelToCurrentGeoSuspicious:   false,
					TravelFromCurrentGeoSuspicious: false,
				},
			},
			expectedStatusCode: http.StatusOK,
			expectedResponse:   `{"currentGeo":{"lat":1,"lon":2,"radius":0},"travelToCurrentGeoSuspicious":false,"travelFromCurrentGeoSuspicious":false}`,
		},
		{
			name:               "bad request unmarshaling",
			payload:            `{sdlkfjsdlfkjsdf}`,
			expectedStatusCode: http.StatusBadRequest,
			expectedResponse:   `{"error": "error decoding request body"}`,
		},
		{
			name: "error",
			payload: `{
				"username":"bob",
				"unix_timestamp":1514764800,
				"event_uuid":"85ad929a-db03-4bf4-9541-8f728fa12e42",
				"ip_address":"206.81.252.6"
			 }`,
			err: errors.New("some error"),
			svc: &fakeDetection{
				err: errors.New("some error"),
			},
			expectedStatusCode: http.StatusInternalServerError,
			expectedResponse:   `{"error": "some error"}`,
		},
	}

	for _, c := range cases {
		dh := &DetectHandler{
			DetectSvc: c.svc,
		}

		r, err := http.NewRequest("POST", "/detect", strings.NewReader(c.payload))
		if err != nil {
			t.Fatal(err)
		}

		w := httptest.NewRecorder()

		dh.Detect(w, r)
		if c.expectedStatusCode != w.Code {
			t.Errorf("TestDetect - %s, expected %d, got %d", c.name, c.expectedStatusCode, w.Code)
		}

		if len(w.Body.String()) != len(c.expectedResponse) {
			t.Errorf("TestDetect - %s, \nexpected %s\ngot %v %d %d", c.name, c.expectedResponse, w.Body.String(), len(w.Body.String()), len(c.expectedResponse))
		}
	}
}
