package workflow

import (
	"github.com/mgranderath/dns-measurements/db"
	"github.com/miekg/dns"
	"time"
)

type workflow struct {
	Message *dns.Msg
	IP string
	RTT *time.Duration
}

func StartForIP(ip string) {
	rtt, err := getRTT(ip)
	if err != nil {
		return
	}
	db.CreateServer(ip, rtt)

	req := dns.Msg{}
	req.Id = dns.Id()
	req.RecursionDesired = true
	req.Question = []dns.Question{
		{Name: "test.com" + ".", Qtype: dns.TypeA, Qclass: dns.ClassINET},
	}

	workfl := &workflow{
		IP: ip,
		RTT: rtt,
		Message: &req,
	}

	workfl.testUDP()
	workfl.testTCP()
	workfl.testTLS()
	workfl.testHTTPS()
	workfl.testQuic()
}
