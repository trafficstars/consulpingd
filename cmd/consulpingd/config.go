package main

import (
	"log"
	"time"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	SourceName             string
	PingInterval           time.Duration `default:"1s"`
	PingTimeout            time.Duration `default:"1m"`
	PrometheusExporterBind string        `default:":55634"`
	Consuls                []string
	RegisterAtConsuls      bool
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
