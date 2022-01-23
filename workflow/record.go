package workflow

import (
	"github.com/mgranderath/dns-measurements/db"
	"github.com/mgranderath/dns-measurements/model"
	"github.com/mgranderath/dnsperf/clients"
	"github.com/mgranderath/dnsperf/metrics"
	"github.com/miekg/dns"
	"github.com/rs/xid"
	"log"
	"time"
)

func (w *workflow) convertToMeasurement(id string, protocol string, result *metrics.Result, response *dns.Msg, cacheWarming bool, err error) model.DNSMeasurement {
	var responseIP *string
	var responseTTL *uint32
	var rCode *int
	var errorString *string

	if response != nil {
		rCode = &response.Rcode
		if len(response.Answer) > 0 {
			answer := response.Answer[0]
			switch record := answer.(type) {
			case *dns.A:
				responseIPstring := record.A.String()
				responseIP = &responseIPstring
				responseTTL = &record.Hdr.Ttl
			}
		}
	}

	if err != nil {
		errStr := err.Error()
		errorString = &errStr
	}

	return model.DNSMeasurement{
		ID:                     id,
		IP:                     w.IP,
		Port:                   w.Port,
		CacheWarming:           cacheWarming,
		UDPSocketSetupDuration: result.UDPSocketSetupDuration,
		TCPHandshakeDuration:   result.TCPHandshakeDuration,
		TLSHandshakeDuration:   result.TLSHandshakeDuration,
		TLSVersion:             result.TLSVersion,
		TLSError:               result.TLSError,
		QUICHandshakeDuration:  result.QUICHandshakeDuration,
		QUICVersion:            result.QUICVersion,
		QUICNegotiatedProtocol: result.QUICNegotiatedProtocol,
		QUICError:              result.QUICError,
		HTTPVersion:            result.HTTPVersion,
		QueryTime:              result.QueryTime,
		TotalTime:              result.TotalTime,
		RCode:                  rCode,
		ResponseIP:             responseIP,
		ResponseTTL:            responseTTL,
		Protocol:               protocol,
		Error:                  errorString,
	}
}

func (w *workflow) runMeasurementAndRecord(protocol string, address string, options clients.Options, id xid.ID, cacheWarming bool) uint64 {
	u, err := clients.AddressToClient(protocol+"://"+address, options)
	if err != nil {
		log.Fatalf("Cannot create an upstream: %s", err)
	}

	reply := u.Exchange(w.Message)

	db.AddMeasurement(w.convertToMeasurement(id.String(), protocol, reply.GetMetrics(), reply.GetResponse(), cacheWarming, reply.GetError()))
	if reply.GetMetrics() != nil && reply.GetMetrics().QLogMessages != nil && !cacheWarming {
		db.AddQLogOutput(id.String(), reply.GetMetrics().QLogMessages)
	}

	quicVersion := uint64(0)
	if protocol == "quic" && reply.GetMetrics() != nil && reply.GetError() == nil {
		quicVersion = *reply.GetMetrics().QUICVersion
	}

	if reply.GetError() != nil {
		log.Printf("ERROR: %s://%s:%d - %s", protocol, w.IP, w.Port, reply.GetError().Error())

	} else {
		log.Printf("SUCCESS: %s://%s:%d ", protocol, w.IP, w.Port)
		ttlChan := make(chan struct{}, 1)
		go func() {
			w.tracert(id.String())
			close(ttlChan)
		}()

		select {
		case <-ttlChan:
			return quicVersion
		case <-time.After(options.Timeout):
			log.Printf("Traceroute timeout for %s://%s:%d", protocol, w.IP, w.Port)
			return quicVersion
		}
	}
	return quicVersion
}
