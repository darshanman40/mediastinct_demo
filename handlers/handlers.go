package handlers

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/darshanman40/mediastinct_demo/services"
)

var (
	httpErr   error
	indexPage = []byte(`<h2>Hello World</h2>`)
)

type page struct {
	title string
	body  []byte
}

//IndexHandler ...
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	p := &page{title: "Welcome", body: indexPage}
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.title, p.body)
}

//GetMeAdHandler ...
func GetMeAdHandler(w http.ResponseWriter, r *http.Request) {
	var getMeAd = services.NewGetMeAdService(Mock)
	switch r.Method {
	case "POST":
		getMeAd.DoPost(w, r)
	case "GET":
		getMeAd.DoGet(w, r)
	case "PUT":
		getMeAd.DoPut(w, r)
	case "DELETE":
		getMeAd.DoDelete(w, r)

	}
}

// RecoverHandler ....
func RecoverHandler(inner http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var err error
		defer func() {
			r := recover()
			if r != nil {
				switch t := r.(type) {
				case string:
					err = errors.New(t)
				case error:
					err = t
				default:
					err = errors.New("Unknown error")
				}

				log.Println("ERR: " + err.Error())
				http.Error(w, "Internal server error", http.StatusInternalServerError)
			}
		}()
		inner.ServeHTTP(w, r)
	})
}
