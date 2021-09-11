package workflow

import (
	"crypto/tls"
	"fmt"
	"github.com/mgranderath/dnsperf/clients"
	"github.com/rs/xid"
	"time"
)

func convertToIpWithPort(w *workflow) string {
	return fmt.Sprintf("%s:%d", w.IP, w.Port)
}

const timeout = time.Millisecond * 1500

func (w *workflow) testUDP() {

	opts := clients.Options{
		Timeout: timeout,
	}

	id := xid.New()

	w.runMeasurementAndRecord("udp", convertToIpWithPort(w), opts, id, true)
	w.runMeasurementAndRecord("udp", convertToIpWithPort(w), opts, id, false)
}

func (w *workflow) testTCP() {

	opts := clients.Options{
		Timeout: timeout,
	}

	id := xid.New()

	w.runMeasurementAndRecord("tcp", convertToIpWithPort(w), opts, id, true)
	w.runMeasurementAndRecord("tcp", convertToIpWithPort(w), opts, id, false)
}

func (w *workflow) testTLS() {

	opts := clients.Options{
		Timeout: timeout,
		TLSOptions: &clients.TLSOptions{
			MinVersion:         tls.VersionTLS10,
			MaxVersion:         tls.VersionTLS13,
			InsecureSkipVerify: true,
			SkipCommonName:     true,
		},
	}

	id := xid.New()

	w.runMeasurementAndRecord("tls", convertToIpWithPort(w), opts, id, true)
	w.runMeasurementAndRecord("tls", convertToIpWithPort(w), opts, id, false)
}

func (w *workflow) testHTTPS() {

	opts := clients.Options{
		Timeout: timeout,
		TLSOptions: &clients.TLSOptions{
			MinVersion:         tls.VersionTLS10,
			MaxVersion:         tls.VersionTLS13,
			InsecureSkipVerify: true,
			SkipCommonName:     true,
		},
	}

	id := xid.New()

	w.runMeasurementAndRecord("https", convertToIpWithPort(w) + "/dns-query", opts, id, true)
	w.runMeasurementAndRecord("https", convertToIpWithPort(w) + "/dns-query", opts, id, false)
}

func (w *workflow) testQuic() {

	opts := clients.Options{
		Timeout: timeout,
		TLSOptions: &clients.TLSOptions{
			MinVersion:         tls.VersionTLS10,
			MaxVersion:         tls.VersionTLS13,
			InsecureSkipVerify: true,
			SkipCommonName:     true,
		},
		QuicOptions: &clients.QuicOptions{},
	}

	id := xid.New()

	w.runMeasurementAndRecord("quic", convertToIpWithPort(w), opts, id, true)
	w.runMeasurementAndRecord("quic", convertToIpWithPort(w), opts, id, false)
}