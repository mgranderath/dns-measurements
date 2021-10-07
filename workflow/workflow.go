package workflow

import (
	"github.com/miekg/dns"
)

type workflow struct {
	Message *dns.Msg
	IP string
	Port int
	Protocol string
}

func Standard(ip string) {
	Start(ip, 53, "udp")
	Start(ip, 53, "tcp")
	Start(ip, 853, "tls")
	Start(ip, 443, "https")
	Start(ip, 784, "quic")
	Start(ip, 8853, "quic")
	Start(ip, 853, "quic")
}

func Start(ip string, port int, protocol string) {
	req := dns.Msg{}
	req.Id = dns.Id()
	req.RecursionDesired = true
	req.Question = []dns.Question{
		{Name: "test.com" + ".", Qtype: dns.TypeA, Qclass: dns.ClassINET},
	}

	workfl := &workflow{
		IP: ip,
		Message: &req,
		Port: port,
		Protocol: protocol,
	}

	switch protocol {
	case "udp":
		workfl.testUDP()
	case "tcp":
		workfl.testTCP()
	case "tls":
		workfl.testTLS()
	case "https":
		workfl.testHTTPS()
	case "quic":
		workfl.testQuic()
	}
}
