package server

import (
	"net/http"

	"github.com/gorilla/mux"
)

func NewHttpServer(addr string) *http.Server {
	httpsrv := newHttpServer()
	r := mux.NewRouter()
	r.HandleFunc("/", httpsrv.handleProduce).Methods("POST")
	r.HandleFunc("/", httpsrv.handleConsume).Methods("GET")
	return &http.Server{
		Addr:    addr,
		Handler: r}
}
