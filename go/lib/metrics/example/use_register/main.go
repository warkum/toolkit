package main

import (
	"context"
	"fmt"
	"github.com/warkum/toolkit/go/lib/metrics"
	"math/rand"
	"net/http"
	"time"

	"go.opentelemetry.io/otel/exporters/prometheus"
	"go.opentelemetry.io/otel/metric/instrument"
)

const myCounter = "test.my"

func main() {
	ctx := context.Background()

	setHttpHandler()

	for {
		n := rand.Intn(1000)
		time.Sleep(time.Duration(n) * time.Millisecond)

		metrics.SetMetrics(ctx, myCounter, metrics.TypeCounter, 1)
	}
}

func setHttpHandler() {
	// init metrics
	exporter, err := metrics.Init("example", prometheus.Config{})
	if err != nil {
		panic(err)
	}

	// register metrics
	err = metrics.RegisterMetrics(myCounter, metrics.TypeCounter, instrument.WithDescription("Just a test counter"))
	if err != nil {
		panic(err)
	}

	// register exporter to http
	http.HandleFunc("/metrics", exporter.ServeHTTP)
	fmt.Println("listenening on http://localhost:8088/metrics")

	go func() {
		_ = http.ListenAndServe(":8088", nil)
	}()
}
