```sh
# install
go get github.com/trafficstars/consulpingd/cmd/consulpingd
go install github.com/trafficstars/consulpingd/cmd/consulpingd

# configure
export CONSULPINGD_PING_TIMEOUT=10s
export CONSULPINGD_SOURCE_NAME=192.168.100.1
export CONSULPINGD_PROMETHEUS_EXPORTER_BIND=:8083
export CONSULPINGD_CONSULS=consul.service.consul:8500/dc1,consul.service.consul:8500/dc2

# run
sudo setcap cap_net_raw+ep `go env GOPATH`/bin/consulpingd
`go env GOPATH`/bin/consulpingd
```

```sh
# get
curl -s http://localhost:8083/
...
metrics_latency_1h_per90{destination="192.168.100.152",service="consulpingd",source="192.168.100.1"} 1.21586e+08
metrics_latency_1h_per90{destination="192.168.100.48",service="consulpingd",source="192.168.100.1"} 1.19018e+08
metrics_latency_1h_per90{destination="192.168.100.114",service="consulpingd",source="192.168.100.1"} 1.17681e+08
metrics_latency_1h_per90{destination="192.168.100.47",service="consulpingd",source="192.168.100.1"} 1.1897e+08
metrics_latency_1h_per90{destination="192.168.100.32",service="consulpingd",source="192.168.100.1"} 1.09545e+08
metrics_latency_1h_per90{destination="192.168.100.40",service="consulpingd",source="192.168.100.1"} 1.04824e+08
...
```
