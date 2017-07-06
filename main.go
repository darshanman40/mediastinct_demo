package main

import (
	"flag"
	"net/http"

	"github.com/darshanman40/mediastinct_demo/handlers"
)

const port = ":8081"

//mock ...
var mock bool

func main() {

	handlers.InitHandlers(mock)
	http.ListenAndServe(port, nil)

}

func init() {
	flag.BoolVar(&mock, "mock", false, "use mock server")
	flag.Parse()
}
