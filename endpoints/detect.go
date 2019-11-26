package endpoints

import (
	"encoding/json"
	"net/http"

	"github.com/dnguy078/go-detector/models"
)

type DetectHandler struct {
	DetectSvc DetectionService
}

type DetectionService interface {
	Suspicious(models.Request) (*models.Response, error)
}

func (rh *DetectHandler) Detect(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, marshalError("method is not supported"), http.StatusBadRequest)
		return
	}

	lm := models.Request{}
	if err := json.NewDecoder(r.Body).Decode(&lm); err != nil {
		http.Error(w, marshalError("error decoding request body"), http.StatusBadRequest)
		return
	}

	resp, err := rh.DetectSvc.Suspicious(lm)
	if err != nil {
		http.Error(w, marshalError(err.Error()), http.StatusInternalServerError)
		return
	}

	b, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, marshalError(err.Error()), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func marshalError(errString string) string {
	type endpointError struct {
		Message string `json:"error"`
	}
	e := &endpointError{Message: errString}
	b, err := json.Marshal(e)
	if err != nil {
		return ""
	}
	return string(b)
}
