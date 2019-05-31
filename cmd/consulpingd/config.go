package main

import (
	"log"
	"time"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	SourceName             string        `envconfig:"SOURCE_NAME"`
	PingInterval           time.Duration `default:"1s"     envconfig:"PING_INTERVAL"`
	PingTimeout            time.Duration `default:"1m"     envconfig:"PING_TIMEOUT"`
	PrometheusExporterBind string        `default:":55634" envconfig:"PROMETHEUS_EXPORTER_BIND"`
	Consuls                []string
	RegisterAtConsuls      bool `envconfig:"REGISTER_AT_CONSULS"`
	Verbose                bool
}

func GetConfig() *Config {
	cfg := &Config{}
	err := envconfig.Process("CONSULPINGD", cfg)
	if err != nil {
		log.Panic(err)
	}
	if cfg.RegisterAtConsuls {
		log.Panic(`Not implemented, yet`)
	}
	return cfg
}
