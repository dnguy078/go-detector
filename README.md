# go-detector
go-detector is an API built to detect possible compromised account credentials based on source IP addresses

## Running locally
Github does not allow users to upload over 60mb files (zipped contains seed data for geoIP, must be ran once)
To Run:
```
unzip ./data/city_seed_data.mmdb.zip -d ./data
go mod download
go run main.go
```
To Test:
```
go test ./...
```
Integration Test:
```
go test ./... -tags=integration
```
Build and Run using docker:
```
docker build -t secureworks .
docker run -p 3000:3000 secureworks
```

### Package dependencies
1. github.com/gchaincl/dotsql v1.0.0 (used to source .sql file)
2. github.com/google/uuid v1.1.1 (used to test with random uuid)
3. github.com/mattn/go-sqlite3 v1.13.0 (used for sqllite)
4. github.com/oschwald/geoip2-golang v1.3 (used to mindmasterdb wrapper)

## API:
### Detect
```json

Request:
TimeStamp: 01/01/2018 @ 12:00am

curl -X POST \
  http://localhost:3000/detect \
  -H 'Content-Type: application/json' \
  -d '{
        "username":  "test",
		"unix_timestamp":  1574735184,
		"event_uuid":  "testing",
		"ip_address":  "107.77.66.81"
    }'

Response:

Preceding: 01/01/2018 @ 12:00am
Subsequent: 01/02/2018 @ 12:00am (UTC)
Speed is calculated by taking distance between locations in miles / diff in time in hours
{
	"currentGeo": {
		"lat": ​39.1702​,
		"lon": ​-76.8538​,
		"radius": 2​0
	},
	"travelToCurrentGeoSuspicious": true,
	"travelFromCurrentGeoSuspicious": false,
	"precedingIpAccess": {
		"ip": "24.242.71.20"​,
		"speed": 1341,
		"lat": ​30.3764​,
		"lon": ​-97.7078​,
		"radius": 5​​,
		"timestamp": 1514764800
	},
	"subsequentIpAccess": {
		"ip": "91.207.175.104"​,
		"speed": 101,
		"lat": ​34.0494​,
		"lon": ​-118.2641​,
		"radius": 2​00​,
		"timestamp": 1514851200
	}
}
```

### Outstanding questions:
1. Did not use the radius, unsure if we should be adding/subtracting to reach location before determing the distance. (speed is not accurate in specs in either case)
2. Speed calculation was not documented in the specs, I set it to distance from two locations in miles / time difference between locations in hours.
