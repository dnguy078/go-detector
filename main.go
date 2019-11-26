package main

import (
	"log"

	"github.com/dnguy078/go-detector/daemon"
)

const (
	port         = ":3000"
	dbPath       = "./geo.db"
	dbSchemaPath = "./schema/geo.db.sql"
	geoSeedPath  = "./data/city_seed_data.mmdb"
)

func main() {
	s, err := daemon.New(port, dbPath, dbSchemaPath, geoSeedPath)
	if err != nil {
		log.Fatal(err)
	}
	if err := s.Start(); err != nil {
		log.Fatalf("unable to start server: %s", err)
	}
}
