package main

import (
	"fmt"
	"log"
	"net"
	"time"

	"github.com/trafficstars/consulpingd/consul"
	"github.com/trafficstars/consulpingd/exporter"
	"github.com/trafficstars/consulpingd/ping"
	"github.com/trafficstars/metrics"
)

func doPing(ip *net.IP, timeout time.Duration) {
	tags := metrics.Tags{
		`destination`: ip.String(),
	}
	result, err := ping.ICMPEcho(ip, timeout)
	if err != nil {
		metrics.GaugeAggregativeBuffered(`errored`, tags).ConsiderValue(1)
		log.Print("unable to ICMP-echo-ping", ip.String(), ":", err)
		return
	}
	metrics.GaugeAggregativeBuffered(`errored`, tags).ConsiderValue(0)
	if result.TimedOut {
		metrics.GaugeAggregativeBuffered(`loss`, tags).ConsiderValue(1)
		return
	}
	metrics.GaugeAggregativeBuffered(`loss`, tags).ConsiderValue(0)
	metrics.TimingBuffered(`latency`, tags).ConsiderValue(result.Latency)
	/*corruptedValue := float64(0)
	if result.Corrupted {
		corruptedValue = 1
	}
	metrics.GaugeAggregativeBuffered(`corrupted`, tags).ConsiderValue(corruptedValue)*/
}

func main() {
	cfg := GetConfig()
	if cfg.Verbose {
		fmt.Println("config:", *cfg)
	}

	metrics.SetDefaultTags(metrics.Tags{
		`service`: `consulpingd`,
		`source`:  cfg.SourceName,
	})

	consulInstance, err := consul.New(cfg.Consuls, cfg.RegisterAtConsuls)
	if err != nil {
		log.Panic(err)
	}
	go func() {
		ticker := time.NewTicker(cfg.PingInterval)
		for {
			go func() {
				ips := consulInstance.GetIPs()
				for _, ip := range ips {
					go doPing(ip, cfg.PingTimeout)
				}
			}()
			<-ticker.C
		}
	}()

	prometheusExporter := exporter.NewPrometheus(cfg.PrometheusExporterBind)
	log.Panic(prometheusExporter.Loop())
}
