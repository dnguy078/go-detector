-- name: load
BEGIN TRANSACTION;
DROP TABLE IF EXISTS "user_geo_events";
CREATE TABLE IF NOT EXISTS "user_geo_events" (
	"username"	TEXT NOT NULL,
	"event_uuid"	TEXT NOT NULL UNIQUE,
	"inserted_at" INTEGER NOT NULL,
	"timestamp"	INTEGER NOT NULL,
	"ip_address"	INTEGER NOT NULL,
	"lat"	REAL NOT NULL,
	"lon"	REAL NOT NULL,
	"radius"	REAL NOT NULL,
	PRIMARY KEY("username", "event_uuid")
);
DROP INDEX IF EXISTS "user_events_index";
CREATE INDEX IF NOT EXISTS "user_events_index" ON "user_geo_events" (
	"username"	ASC,
	"timestamp"	ASC
);
COMMIT;

