```sh
# install
go get github.com/trafficstars/consulpingd/cmd/consulpingd
go install github.com/trafficstars/consulpingd/cmd/consulpingd

# configure
export CONSULPINGD_PING_TIMEOUT=10s
export CONSULPINGD_SOURE_NAME=10.0.0.1
export CONSULPINGD_PROMETHEUS_EXPORTER_BIND=:8083
export CONSULPINGD_CONSULS=consul.service.consul:8500

# run
`go env GOPATH`/bin/consulpingd
```

```sh
# get
curl http://localhost:8083/
```