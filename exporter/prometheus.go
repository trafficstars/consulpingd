package exporter

import (
	"log"
	"net/http"

	"github.com/trafficstars/statuspage"
)

type Prometheus struct {
	bindAddress string
}

func NewPrometheus(bindAddress string) *Prometheus {
	return &Prometheus{
		bindAddress: bindAddress,
	}
}

func (p *Prometheus) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := statuspage.WriteMetricsPrometheus(w)
	if err != nil {
		log.Print(err)
	}
}

func (p *Prometheus) Loop() error {
	srv := &http.Server{
		Addr:           p.bindAddress,
		Handler:        p,
		MaxHeaderBytes: 1 << 20,
	}
	return srv.ListenAndServe()
}
