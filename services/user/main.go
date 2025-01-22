package main

import (
	"net/http"

	userservice "github.com/kasyap1234/book-review/services/user/internal"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	go func() {
        http.Handle("/metrics", promhttp.Handler())
        http.ListenAndServe("0.0.0.0:2112", nil)
    }()

	userservice.StartServer3()

}
