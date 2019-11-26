package adapter

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/dnguy078/go-detector/models"
	"github.com/gchaincl/dotsql"
	_ "github.com/mattn/go-sqlite3"
)

// db holds a database connection to a sqllite database
type db struct {
	conn *sql.DB
}

func NewDB(dbLocationPath string) (*db, error) {
	conn, err := sql.Open("sqlite3", dbLocationPath)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return &db{
		conn: conn,
	}, nil
}

func (d *db) Close() error {
	return d.conn.Close()
}

// LoadCSV loads a csv file into a table, sqllite3 package does not have feature to `Load` from file
func (d *db) LoadFromFile(path string) error {
	dot, err := dotsql.LoadFromFile(path)
	if err != nil {
		return err
	}
	if _, err := dot.Exec(d.conn, "load"); err != nil {
		return err
	}

	return nil
}

// InsertUserEvent inserts entry
func (d *db) InsertUserEvent(e *models.UserGeoEvent) error {
	statement, err := d.conn.Prepare(`INSERT INTO user_geo_events(
		username,
		event_uuid,
		inserted_at,
		timestamp,
		ip_address,
		lat,
		lon,
		radius) VALUES (?,?,?,?,?,?,?,?)`)
	if err != nil {
		return err
	}
	_, err = statement.Exec(e.Username, e.EventUUID, time.Now().Unix(), e.Timestamp, e.IPAddress, e.Lat, e.Lon, e.Radius)

	return err
}

// GetPreviousIpAccess returns the last previous location traveled
func (d *db) GetPreviousIPAccess(e *models.UserGeoEvent) (*models.UserGeoEvent, error) {
	query := `SELECT *
            FROM user_geo_events
            WHERE username=? AND timestamp <= ? and event_uuid != ?
            ORDER BY timestamp DESC
			LIMIT 1
	`
	prev := &models.UserGeoEvent{
		Location: new(models.Location),
	}
	err := d.conn.QueryRowContext(context.Background(), query, e.Username, e.Timestamp, e.EventUUID).Scan(
		&prev.Username, &prev.EventUUID, &prev.InsertedAt, &prev.Timestamp, &prev.IPAddress, &prev.Lat, &prev.Lon, &prev.Radius)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return nil, nil
		}
		return nil, err
	}

	return prev, nil
}

// GetSubsequentIpAccess returns the subsequent location traveled
func (d *db) GetSubsequentIPAccess(e *models.UserGeoEvent) (*models.UserGeoEvent, error) {
	query := `SELECT *
            FROM user_geo_events
            WHERE username=? AND timestamp > ?
            ORDER BY inserted_at ASC
			LIMIT 1
	`
	subsequent := &models.UserGeoEvent{
		Location: new(models.Location),
	}

	err := d.conn.QueryRowContext(context.Background(), query, e.Username, e.Timestamp, e.EventUUID).Scan(
		&subsequent.Username,
		&subsequent.EventUUID,
		&subsequent.InsertedAt,
		&subsequent.Timestamp,
		&subsequent.IPAddress,
		&subsequent.Lat,
		&subsequent.Lon,
		&subsequent.Radius,
	)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return nil, nil
		}
		return nil, err
	}

	return subsequent, nil
}
