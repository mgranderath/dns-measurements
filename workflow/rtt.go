package workflow

import (
	"errors"
	"github.com/go-ping/ping"
	"github.com/mgranderath/dns-measurements/db"
	"github.com/mgranderath/dns-measurements/model"
	"github.com/mgranderath/traceroute/methods"
	"github.com/mgranderath/traceroute/methods/tcp"
	"github.com/mgranderath/traceroute/methods/udp"
	"log"
	"net"
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

func (w *workflow) tracert(id string) {
	network := "udp"

	switch w.Protocol {
	case "tcp":
		fallthrough
	case "tls":
		fallthrough
	case "https":
		network = "tcp"
	}

	if network == "udp" {
		udpTraceroute := udp.New(net.ParseIP(w.IP), w.Protocol == "quic", methods.TracerouteConfig{
			MaxHops:          30,
			NumMeasurements:  5,
			ParallelRequests: 15,
			Port:             w.Port,
			Timeout:          time.Second,
		})

		result, err := udpTraceroute.Start()
		if err != nil || result == nil {
			log.Println(err)
		}

		for key, hopResults := range *result {
			for _, hop := range hopResults {
				if !hop.Success {
					db.AddTraceroute(model.Traceroute{
						DNSMeasurementID: id,
						Timestamp:        time.Now(),
						TTL:              key,
						DestPort: w.Port,
						Protocol: w.Protocol,
						DestIP: w.IP,
						HopIP:            nil,
						RTT:              nil,
					})
					continue
				}
				hopAddress := hop.Address.String()
				db.AddTraceroute(model.Traceroute{
					DNSMeasurementID: id,
					Timestamp:        time.Now(),
					DestPort: w.Port,
					Protocol: w.Protocol,
					TTL:              key,
					HopIP:            &hopAddress,
					RTT:              hop.RTT,
					DestIP: w.IP,
				})
			}
		}
	} else {
		tcpTraceroute := tcp.New(net.ParseIP(w.IP), methods.TracerouteConfig{
			MaxHops:          30,
			NumMeasurements:  5,
			ParallelRequests: 15,
			Port:             w.Port,
			Timeout:          time.Second,
		})

		result, err := tcpTraceroute.Start()
		if err != nil || result == nil {
			log.Println(err)
		}

		for key, hopResults := range *result {
			for _, hop := range hopResults {
				if !hop.Success {
					db.AddTraceroute(model.Traceroute{
						DNSMeasurementID: id,
						Timestamp:        time.Now(),
						DestPort: w.Port,
						Protocol: w.Protocol,
						TTL:              key,
						HopIP:            nil,
						RTT:              nil,
						DestIP: w.IP,
					})
					continue
				}
				hopAddress := hop.Address.String()
				db.AddTraceroute(model.Traceroute{
					DNSMeasurementID: id,
					Timestamp:        time.Now(),
					DestPort: w.Port,
					Protocol: w.Protocol,
					TTL:              key,
					HopIP:            &hopAddress,
					RTT:              hop.RTT,
					DestIP: w.IP,
				})
			}
		}
	}
}
