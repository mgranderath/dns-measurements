package workflow

import (
	"crypto/tls"
	"github.com/mgranderath/dnsperf/clients"
)

func (w *workflow) testUDP() {
	timeout := 50 * *w.RTT

	opts := clients.Options{
		Timeout: timeout,
	}

	w.runMeasurementAndRecord("udp", w.IP, opts)
}

func (w *workflow) testTCP() {
	timeout := 50 * *w.RTT

	opts := clients.Options{
		Timeout: timeout,
	}

	w.runMeasurementAndRecord("tcp", w.IP, opts)
}

func (w *workflow) testTLS() {
	timeout := 50 * *w.RTT

	opts := clients.Options{
		Timeout: timeout,
		TLSOptions: &clients.TLSOptions{
			MinVersion:         tls.VersionTLS10,
			MaxVersion:         tls.VersionTLS13,
			InsecureSkipVerify: true,
			SkipCommonName:     true,
		},
	}

	w.runMeasurementAndRecord("tls", w.IP, opts)
}

func (w *workflow) testHTTPS() {
	timeout := 50 * *w.RTT

	opts := clients.Options{
		Timeout: timeout,
		TLSOptions: &clients.TLSOptions{
			MinVersion:         tls.VersionTLS10,
			MaxVersion:         tls.VersionTLS13,
			InsecureSkipVerify: true,
			SkipCommonName:     true,
		},
	}

	w.runMeasurementAndRecord("https", w.IP + "/dns-query", opts)
}

func (w *workflow) testQuic() {
	timeout := 50 * *w.RTT

	opts := clients.Options{
		Timeout: timeout,
		TLSOptions: &clients.TLSOptions{
			MinVersion:         tls.VersionTLS10,
			MaxVersion:         tls.VersionTLS13,
			InsecureSkipVerify: true,
			SkipCommonName:     true,
		},
		QuicOptions: &clients.QuicOptions{
			AllowedVersions: []string{clients.VersionQuic00, clients.VersionQuic01, clients.VersionQuic02},
		},
	}

	w.runMeasurementAndRecord("quic", w.IP, opts)
}