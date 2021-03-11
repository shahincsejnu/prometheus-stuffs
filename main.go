package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/gorilla/mux"
)

func home(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello from demo application"))
}

func notfound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("not found"))
}

var (
	version = prometheus.NewGauge(prometheus.GaugeOpts{
		Name:        "version",
		Help:        "Version information about the binary",
		ConstLabels: map[string]string{
			"version": "v1.0.1",
		},
	})

	httpRequestsTotal = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name:        "http_requests_total",
		Help:        "Count of all HTTP requests",
	}, []string{"code", "method"})
)

func main() {
	bind := ""
	flagset := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	flagset.StringVar(&bind, "bind", ":8080", "The socket to bind to.")
	flagset.Parse(os.Args[1:])

	r := prometheus.NewRegistry()
	r.MustRegister(httpRequestsTotal)
	r.MustRegister(version)

	handler := mux.NewRouter().StrictSlash(true)

	homeHandler := http.HandlerFunc(home)
	notfoundHandler := http.HandlerFunc(notfound)

	handler.HandleFunc("/", promhttp.InstrumentHandlerCounter(httpRequestsTotal, homeHandler))
	handler.HandleFunc("/err", promhttp.InstrumentHandlerCounter(httpRequestsTotal, notfoundHandler))

	handler.Handle("/metrics", promhttp.HandlerFor(r, promhttp.HandlerOpts{}))

	fmt.Println("Server is listening in 8080 port...")
	log.Fatal(http.ListenAndServe(bind, handler))
}