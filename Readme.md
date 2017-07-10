#Ad API Demo: Async Ad Server

## To run the api: <br>
`go run main.go`

## To run API with mock data: <br>
`go run main.go -mock=true`
### Note: Start the mock service using following:
#### To start Nike mock ad server
`go run mockserver/main.go -profile "mocknikead"`
#### To start Amazon mock ad server
`go run mockserver/main.go -profile "mockamazonad"`
#### To start eBay mock ad server
`go run mockserver/main.go -profile "mockebayad"`

The API runs on port 8081 and mock server runs on 8082. Using postman make Get request on http://localhost:8081/getmead?gender=male&age=20

It will read config of client URLS from `client/client_ads.json` and pick 'adcode' of highest bid. Feel free to add urls, client names in that json file and play around.


##### Bugs:
- ~~Inconsistency with highest biding Ad (06-July-17)~~: Fixed with locking-unlocking sync.Mutex for each client request (07-July-17).
