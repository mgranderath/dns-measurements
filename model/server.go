package model

import "time"

type Server struct {
	IP string `gorm:"primaryKey"`
	Created   int64 `gorm:"autoCreateTime"`
	Updated int64 `gorm:"autoUpdateTime"`
	RTT *time.Duration
	Measurements []Measurement
}
