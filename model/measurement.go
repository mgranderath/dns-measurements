package model

import "time"

type Measurement struct {
	ServerID string
	Created   int64 `gorm:"autoCreateTime"`
	Updated int64 `gorm:"autoUpdateTime"`
	Protocol string

	UDPSocketSetupDuration *time.Duration `json:"udp_socket_setup_duration,omitempty"`
	TCPHandshakeDuration *time.Duration `json:"tcp_handshake_duration,omitempty"`

	TLSHandshakeDuration *time.Duration `json:"tls_handshake_duration,omitempty"`
	TLSVersion           *uint16        `json:"tls_version,omitempty"`
	TLSError             *int           `json:"tls_error,omitempty"`

	QUICHandshakeDuration  *time.Duration `json:"quic_handshake_duration,omitempty"`
	QUICVersion            *uint64        `json:"quic_version,omitempty"`
	QUICNegotiatedProtocol *string        `json:"quic_negotiated_protocol,omitempty"`
	QUICError              *uint64        `json:"quic_error,omitempty"`

	HTTPVersion *string `json:"http_version,omitempty"`

	QueryTime *time.Duration `json:"query_time,omitempty"`

	TotalTime *time.Duration `json:"total_time,omitempty"`

	RCode *int
	ResponseIP *string
	ResponseTTL *uint32

	Error *string
}
