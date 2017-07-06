package services

import "net/http"

//Service ...
type Service interface {
	DoGet(w http.ResponseWriter, r *http.Request)
	DoPost(w http.ResponseWriter, r *http.Request)
	DoPut(w http.ResponseWriter, r *http.Request)
	DoDelete(w http.ResponseWriter, r *http.Request)
}
