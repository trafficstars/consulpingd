package ping

import (
	"time"
)

type Result struct {
	TimedOut  bool
	Corrupted bool
	Latency   time.Duration
}
