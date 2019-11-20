package endpoints

import (
	"encoding/json"
	"net/http"

	"github.com/dnguy078/go-detector/models"
)

type DetectHandler struct {
	Dispatcher Broker
}

type Broker interface {
	Send(models.Loginmeta) error
}

func (rh *DetectHandler) Detect(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method is not supported", http.StatusBadRequest)
		return
	}

	lm := models.Loginmeta{}
	if err := json.NewDecoder(r.Body).Decode(&lm); err != nil {
		http.Error(w, "error decoding request body", http.StatusBadRequest)
		return
	}
	if lm.Username == "" {
		http.Error(w, "requestID cannot not be empty", http.StatusBadRequest)
		return
	}

	// if err := rh.Dispatcher.Send(*record); err != nil {
	// 	http.Error(w, "unable to process message at this time", http.StatusInternalServerError)
	// 	return
	// }

	w.WriteHeader(http.StatusAccepted)
}
