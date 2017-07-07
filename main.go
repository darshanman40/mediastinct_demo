package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/darshanman40/mediastinct_demo/handlers"
)

const portColon = ":"

var port string
var mock bool

func main() {
	initArguments()
	if mock {
		log.Println("============================================================")
		log.Println("============== WARNING: Using Mock data ====================")
		log.Println("====== Make sure mock server is running at port 8082 =======")
		log.Println("============================================================")
	}
	handlers.InitHandlers(mock)
	http.ListenAndServe(portColon+port, nil)

}

func initArguments() {
	flag.BoolVar(&mock, "mock", false, "use mock server")
	flag.StringVar(&port, "port", "80", "server port to be run")
	flag.Parse()
}
