package workflow

import (
	"errors"
	"github.com/go-ping/ping"
	"runtime"
	"time"
)

func getRTT(ip string) (*time.Duration, error) {
	pinger, err := ping.NewPinger(ip)
	if err != nil {
		return nil, err
	}
	if runtime.GOOS == "windows" {
		pinger.SetPrivileged(true)
	}
	pinger.Count = 10
	pinger.Interval = time.Millisecond * 200
	pinger.Timeout = time.Second * 5
	err = pinger.Run()
	if err != nil {
		return nil, err
	}

	stats := pinger.Statistics()
	if stats.PacketsRecv == 0 {
		return nil, errors.New("could not reach host")
	}

	return &stats.AvgRtt, nil
}
