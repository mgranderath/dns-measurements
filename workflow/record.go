package workflow

import (
	"github.com/mgranderath/dns-measurements/db"
	"github.com/mgranderath/dns-measurements/model"
	"github.com/mgranderath/dnsperf/clients"
	"github.com/mgranderath/dnsperf/metrics"
	"github.com/miekg/dns"
	"log"
)

func (w *workflow) convertToMeasurement(protocol string, result *metrics.Result, response *dns.Msg, err error) model.Measurement {
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

	return model.Measurement{
		ServerID:               w.IP,
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
		RCode: rCode,
		ResponseIP: responseIP,
		ResponseTTL: responseTTL,
		Protocol: protocol,
		Error: errorString,
	}
}

func (w *workflow) runMeasurementAndRecord(protocol string, address string, options clients.Options) {
	u, err := clients.AddressToClient(protocol+"://"+address, options)
	if err != nil {
		log.Fatalf("Cannot create an upstream: %s", err)
	}

	reply := u.Exchange(w.Message)

	db.AddMeasurement(w.convertToMeasurement(protocol, reply.GetMetrics(), reply.GetResponse(), reply.GetError()))

	if reply.GetError() != nil {
		log.Printf("%s://%s measurement error: %s", protocol, w.IP, reply.GetError().Error())
	} else {
		log.Printf("%s://%s measurement successfull", protocol, w.IP)
	}
}
