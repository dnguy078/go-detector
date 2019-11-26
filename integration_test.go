// +build integration

package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
	"testing"
	"time"

	"github.com/dnguy078/go-detector/daemon"
	"github.com/dnguy078/go-detector/models"
)

var httpClient = http.Client{}

func TestIntegration(t *testing.T) {
	s, err := daemon.New(":8081", ":memory:", "./schema/geo.db.sql", "./data/city_seed_data.mmdb")
	if err != nil {
		log.Fatal(err)
	}
	go s.Start()
	time.Sleep(1 * time.Second)

	if len(requests) != len(responses) {
		t.Error("request and response should be same length")
	}

	for i, req := range requests {
		b, err := json.Marshal(req)
		if err != nil {
			t.Error(err)
		}
		resp, err := http.Post("http://localhost:8081/detect", "application/json", bytes.NewBuffer(b))
		if err != nil {
			log.Fatalln(err)
		}
		if resp.StatusCode != http.StatusOK {
			t.Error("expected to return 2XX status code")
		}
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			t.Fatal(err)
		}

		actualResponse := models.Response{}
		if err := json.Unmarshal(body, &actualResponse); err != nil {
			t.Error(err)
		}

		if !reflect.DeepEqual(responses[i], actualResponse) {
			t.Errorf("Test Request #%d with IPAddress: %s", i, req.IPAddress)
			t.Errorf("\n%+v - \n%+v", actualResponse, responses[i])
			t.Errorf("Current locationgot \n%+v\nexpected: \n%+v\n", actualResponse.CurrentGeo, responses[i].CurrentGeo)
			t.Errorf("PreviousIP Address got \n%+v\nexpected: \n%+v\n", actualResponse.PrecedingIPAccess, responses[i].PrecedingIPAccess)
			t.Errorf("NextIP Address got \n%+v\nexpected: \n%+v\n", actualResponse.SubsequentIPAccess, responses[i].SubsequentIPAccess)
		}
	}
}
