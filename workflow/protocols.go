package workflow

import (
	"crypto/tls"
	"fmt"
	"github.com/Lucapaulo/dnsperf/clients"
	"github.com/lucas-clemente/quic-go"
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
	tokenStore := quic.NewLRUTokenStore(5, 50)

	opts := clients.Options{
		Timeout: timeout,
		TLSOptions: &clients.TLSOptions{
			MinVersion:         tls.VersionTLS10,
			MaxVersion:         tls.VersionTLS13,
		},
		QuicOptions: &clients.QuicOptions{
			TokenStore: tokenStore,
			QuicVersions: []quic.VersionNumber{quic.VersionDraft34, quic.VersionDraft32, quic.VersionDraft29, quic.Version1},
		},
	}

	id := xid.New()

	quicVersion := w.runMeasurementAndRecord("quic", convertToIpWithPort(w), opts, id, true)
	if quicVersion != uint64(0) {
		opts.QuicOptions.QuicVersions = []quic.VersionNumber{quic.VersionNumber(uint32(quicVersion))}
	}
	w.runMeasurementAndRecord("quic", convertToIpWithPort(w), opts, id, false)
}