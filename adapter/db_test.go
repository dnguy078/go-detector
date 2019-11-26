package adapter

import (
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func Test_db_LoadCSV(t *testing.T) {
	tests := []struct {
		name         string
		testFilepath string
		wantErr      bool
	}{
		{
			name:         "success",
			testFilepath: "../data/geo_ip.csv",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, err := NewDB(":memory:")
			if err != nil {
				t.Fatalf("unable to establish sqllite connection, err: %s", err)
			}
			defer db.Close()

			if err := db.LoadFromFile("../schema/geo.db.sql"); (err != nil) != tt.wantErr {
				t.Errorf("db.LoadCSV() - %s - error = %v, wantErr %v", tt.name, err, tt.wantErr)
			}
		})
	}
}
