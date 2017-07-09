package httpclient

import (
	"log"
	"net/http"
	"sync"
	"time"
)

const twoSecDuration = 2 * time.Second

var mutex *sync.Mutex

//HTTPClient ...
type HTTPClient interface {
	Do(*http.Request) (*http.Response, error)
}

type httpClient struct {
	*http.Client
}

func (h *httpClient) Do(r *http.Request) (*http.Response, error) {
	defer func() {
		log.Println("Doing: " + r.URL.RequestURI())
	}()
	resp, err := h.Client.Do(r)
	return resp, err
}

//NewHTTPClient ...
func NewHTTPClient() HTTPClient {
	client := &http.Client{
		Timeout: twoSecDuration,
	}
	mutex = &sync.Mutex{}
	return &httpClient{client}
}
