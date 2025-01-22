package main

import (
	"net/http"
	bookservice "github.com/kasyap1234/book-review/services/book/internal"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)


func main() {
	go func() {
        http.Handle("/metrics", promhttp.Handler())
        http.ListenAndServe("0.0.0.0:2111", nil)
    }()
	bookservice.StartServer1()
	 



}
