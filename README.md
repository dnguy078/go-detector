# go-detector
go-detector is an API built to detect possible compromised account credentials based on source IP addresses

## Running locally
To Run:
```
go mod download
docker-compose build && docker-compose up
```
To Test:
```
make test
```

### Package dependencies
1.
2.
3.

## API:
### Detect
```json

Request:
{
   "username":"bob",
   "unix_timestamp":1514764800,
   "event_uuid":"85ad929a-db03-4bf4-9541-8f728fa12e42",
   "ip_address":"206.81.252.6"
}

Response:

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
		"speed": 55,
		"lat": ​30.3764​,
		"lon": ​-97.7078​,
		"radius": 5​​,
		"timestamp": 1514764800
	},
	"subsequentIpAccess": {
		"ip": "91.207.175.104"​,
		"speed": 27600,
		"lat": ​34.0494​,
		"lon": ​-118.2641​,
		"radius": 2​00​,
		"timestamp": 1514851200
	}
}
```
