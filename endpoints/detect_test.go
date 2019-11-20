package endpoints

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	// "golang.org/x/time/rate"
)

// type fakeHandler struct {
// 	mu     sync.Mutex
// 	called int
// }

// func (fh *fakeHandler) Handler(w http.ResponseWriter, r *http.Request) {
// 	fh.mu.Lock()
// 	defer fh.mu.Unlock()
// 	fh.called++
// }

// type fakeBroker struct {
// 	err error
// }

// func (f *fakeBroker) Send(shared.Record) error {
// 	return f.err
// }

func TestRecord(t *testing.T) {
	cases := []struct {
		name    string
		payload string
		err     error

		expectedStatusCode int
		expectedBody       string
	}{
		// {
		// 	name:               "success",
		// 	payload:            `{"id":"test product","billing": {"name": "somename", "address": "some billing adressw"} }`,
		// 	expectedStatusCode: http.StatusAccepted,
		// },
		// {
		// 	name:               "bad request unmarshalling",
		// 	payload:            `{"id":"test product","billing": {"name": "somename", "address": "some billing adressw"}`,
		// 	expectedStatusCode: http.StatusBadRequest,
		// 	expectedBody:       "error decoding request body\n",
		// },
		// {
		// 	name:               "error",
		// 	payload:            `{"id":"test product","billing": {"name": "somename", "address": "some billing adressw"} }`,
		// 	err:                errors.New("some error"),
		// 	expectedStatusCode: http.StatusInternalServerError,
		// 	expectedBody:       "unable to process message at this time\n",
		// },
	}

	for _, c := range cases {
		dh := &DetectHandler{
			// Dispatcher: fb,
		}

		r, err := http.NewRequest("POST", "/detect", strings.NewReader(c.payload))
		if err != nil {
			t.Fatal(err)
		}
		w := httptest.NewRecorder()

		dh.Detect(w, r)
		if c.expectedStatusCode != w.Code {
			t.Errorf("test - %s, expected %d, got %d", c.name, c.expectedStatusCode, w.Code)
		}
		if c.expectedBody != w.Body.String() {
			t.Errorf("test - %s, expected %q, got %q", c.name, c.expectedBody, w.Body.String())
		}
	}
}
