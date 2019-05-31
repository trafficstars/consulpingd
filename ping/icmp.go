package ping

import (
	"github.com/sparrc/go-ping"
	"net"
	"time"
)

func ICMPEcho(ip *net.IP, timeout time.Duration) (*Result, error) {
	pinger, err := ping.NewPinger(ip.String())
	if err != nil {
		return nil, err
	}
	pinger.SetPrivileged(true)
	pinger.Count = 1
	pinger.Timeout = timeout
	pinger.Run()
	stats := pinger.Statistics()
	r := &Result{
		Latency:  stats.AvgRtt,
		TimedOut: stats.PacketLoss > 0,
	}
	return r, nil
}
