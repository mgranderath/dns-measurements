package model

import "time"

type Traceroute struct {
	DNSMeasurementID string
	Timestamp time.Time
	TTL uint16
	DestIP string
	DestPort int
	Protocol string
	HopIP *string
	RTT *time.Duration
}
