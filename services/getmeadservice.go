package services

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"sync"

	"github.com/darshanman40/mediastinct_demo/data"
	"github.com/darshanman40/mediastinct_demo/http_client"
)

const (
	genderTag = "gender"
	ageTag    = "age"

	postMethodNotAllowed   = "Post method not allowed"
	getMethodNotAllowed    = "Post method not allowed"
	putMethodNotAllowed    = "Post method not allowed"
	deleteMethodNotAllowed = "Post method not allowed"
)

var mutex *sync.Mutex

type page struct {
	title string
	body  []byte
}

//RespData ...
type RespData struct {
	Adcode string `json:"adcode"`
}

type getMeAdService struct {
	bids chan *httpclient.RespAd
	// urlStrs    []string
	clientURLs []data.ClientURLs
}

func (g *getMeAdService) DoGet(w http.ResponseWriter, r *http.Request) {
	var gender = r.FormValue(genderTag)
	var age, err = strconv.Atoi(r.FormValue(ageTag))
	if err != nil {
		log.Println("Failed to conver age datatype: " + err.Error())
		http.Error(w, "500", http.StatusInternalServerError)
		return
	}

	rm := httpclient.NewRequestManager(gender, age, g.bids)
	adCode := process(g.clientURLs, rm, g.bids)

	if adCode == "" {
		log.Println("No adCode found")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNoContent)
		// http.Error(w, "500: Internal Server Error", http.StatusInternalServerError)
	} else {
		rd := RespData{Adcode: adCode}
		if data, err := json.Marshal(rd); err != nil {
			log.Println("Marshal failed for adCode")
			http.Error(w, "500: Internal Server Error", http.StatusInternalServerError)
		} else {
			w.Header().Set("Content-Type", "application/json")
			w.Write(data)
		}
	}

}

func (g *getMeAdService) DoPost(w http.ResponseWriter, r *http.Request) {
	http.Error(w, postMethodNotAllowed, http.StatusMethodNotAllowed)
}

func (g *getMeAdService) DoPut(w http.ResponseWriter, r *http.Request) {
	http.Error(w, putMethodNotAllowed, http.StatusMethodNotAllowed)
}

func (g *getMeAdService) DoDelete(w http.ResponseWriter, r *http.Request) {
	http.Error(w, deleteMethodNotAllowed, http.StatusMethodNotAllowed)
}

func process(clientURLs []data.ClientURLs, rm httpclient.RequestManager, rspAd chan *httpclient.RespAd) string {
	var maxBid float64
	var adCode string

	rm.Work(clientURLs)
	works := int32(len(clientURLs))

	for {
		select {
		case bid := <-rspAd:
			mutex.Lock()
			if bid != nil {
				if currentBid, err := strconv.ParseFloat(bid.Bid, 64); err != nil {
					log.Println("Malformed floatpoint: " + bid.Bid)
				} else {
					if maxBid < currentBid {
						maxBid = currentBid
						adCode = bid.AdCode
					}
				}

			}
			// atomic.AddInt32(&works, -1)
			works--
			mutex.Unlock()
		default:
			if works == 0 {
				return adCode
			}
		}
	}
}

//NewGetMeAdService ...
func NewGetMeAdService(mock bool) Service {
	mutex = &sync.Mutex{}
	// var urlStrs []string
	var clientURLs []data.ClientURLs
	if mock {
		clientURLs = []data.ClientURLs{
			data.ClientURLs{
				Name: "Nike",
				URL:  "http://localhost:8082/nike/getad",
			},
			data.ClientURLs{
				Name: "amazon",
				URL:  "http://localhost:8082/amazon/getmead",
			},
			// data.ClientURLs{
			// 	Name: "ebay",
			// 	URL:  "http://localhost:8082/ebay/getmead",
			// },
		}
	} else {
		clientURLs = data.ParseAll()
	}

	return &getMeAdService{bids: make(chan *httpclient.RespAd, len(clientURLs)),
		clientURLs: clientURLs,
	}
}
