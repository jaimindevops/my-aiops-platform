package main

import (
    "context"
    "fmt"
    "log"
    "net/http"
    "os"

    "github.com/redis/go-redis/v9"
    "github.com/prometheus/client_golang/prometheus"
    "github.com/prometheus/client_golang/prometheus/promhttp"
)

var ctx = context.Background()

// 1. Define the Metric (The Sensor)
var visitorCounter = prometheus.NewCounter(
    prometheus.CounterOpts{
        Name: "aiops_visitor_count_total",
        Help: "Total number of visitors to the AIOps platform",
    },
)

func main() {
    // 2. Register the Sensor
    prometheus.MustRegister(visitorCounter)

    redisHost := os.Getenv("REDIS_HOST")
    if redisHost == "" {
        redisHost = "localhost:6379"
    }

    rdb := redis.NewClient(&redis.Options{
        Addr: redisHost,
    })

    // 3. Expose the /metrics endpoint for Prometheus to scrape
    http.Handle("/metrics", promhttp.Handler())

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        // Increment Redis
        val, err := rdb.Incr(ctx, "visitors").Result()
        if err != nil {
            http.Error(w, "Error connecting to Redis", http.StatusInternalServerError)
            return
        }
        
        // 4. Increment the Prometheus Metric (The Signal)
        visitorCounter.Inc()

        msg := fmt.Sprintf("AIOps Platform - Visitor Count: %d", val)
        fmt.Fprintf(w, msg)
    })

    port := ":8080"
    fmt.Printf("Master Node App starting on %s...\n", port)
    if err := http.ListenAndServe(port, nil); err != nil {
        log.Fatal(err)
    }
}
