package main

import (
	"context"
	"fmt"
	"github.com/warkum/toolkit/go/lib/metrics"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"math/rand"
	"net/http"
	"path/filepath"
	"time"
)

const configFile = "./config.yaml"

type Config struct {
	Metric metrics.InitConfig `json:"metric" yaml:"metric"`
}

const (
	testSentMetric = "abc.test_sent"
)

func main() {
	// load config file
	cfg := loadConfig(configFile)
	fmt.Printf("%+v", cfg)

	// set metric exporter based on config
	setMetricsHandler(cfg.Metric)

	ctx := context.Background()

	for {
		n := rand.Intn(1000)
		sleepDur := time.Duration(n) * time.Millisecond
		time.Sleep(sleepDur)

		metrics.SetMetrics(ctx, testSentMetric, metrics.TypeCounter, 1)
		metrics.SetMetrics(ctx, testSentMetric, metrics.TypeHistogram, int64(sleepDur))
	}
}

func loadConfig(configPath string) Config {
	filename, _ := filepath.Abs(configPath)
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	var parsedConfig Config
	err = yaml.Unmarshal(yamlFile, &parsedConfig)
	if err != nil {
		panic(err)
	}

	return parsedConfig
}

func setMetricsHandler(cfg metrics.InitConfig) {
	exporter, err := metrics.InitUsingConfig(cfg)
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/metrics", exporter.ServeHTTP)
	fmt.Println("listenening on http://localhost:8088/metrics")

	go func() {
		_ = http.ListenAndServe(":8088", nil)
	}()
}
