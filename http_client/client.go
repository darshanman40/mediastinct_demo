package httpclient

import (
	"log"
	"net/http"
	"time"
)

const twoSecDuration = 2 * time.Second

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
	return h.Do(r)
}

//NewHTTPClient ...
func NewHTTPClient() HTTPClient {
	client := &http.Client{
		Timeout: twoSecDuration,
	}
	return &httpClient{client}
}
