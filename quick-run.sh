#!/bin/sh


echo "==== Starting main api server at 8080 port ===="
echo "==== Setting mock server configurations ===="
go run main.go -mock=true -port "8080" &

echo "==== Starting mock nike ad server ===="
go run mockserver/main.go -profile "mocknikead" &

echo "==== Starting mock amazon ad server ===="
go run mockserver/main.go -profile "mockamazonad" &

echo "==== Starting mock ebay ad server ===="
go run mockserver/main.go -profile "mockebayad" &
