package metrics

import (
    "net/http"
    "github.com/prometheus/client_golang/prometheus/promhttp"
)

// In your main function or service setup
func StartMetricsServer() {
    
        http.Handle("/metrics", promhttp.Handler())
        http.ListenAndServe(":2111", nil)  // Use 2111 for book, 2112 for user, 2113 for review
 
}
