package main

import (
	"net/http"

	reviewservice "github.com/kasyap1234/book-review/services/review/internal"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	go func() {
        http.Handle("/metrics", promhttp.Handler())
        http.ListenAndServe("0.0.0.0:2113", nil)
    }()
	reviewservice.StartServer2()

}
