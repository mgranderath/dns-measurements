module github.com/mgranderath/dns-measurements

go 1.16

replace github.com/lucas-clemente/quic-go => ./replacement_modules/quic-go

require (
	github.com/Lucapaulo/dnsperf v0.0.8
	github.com/go-ping/ping v0.0.0-20210312085107-d90f3778a8a3
	github.com/jinzhu/now v1.1.2 // indirect
	github.com/lucas-clemente/quic-go v0.21.2
	github.com/mattn/go-sqlite3 v1.14.6 // indirect
	github.com/mgranderath/traceroute v0.0.0-20210421123016-9de04a371a01
	github.com/miekg/dns v1.1.41
	github.com/rs/xid v1.3.0
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c
	gorm.io/driver/sqlite v1.1.4
	gorm.io/gorm v1.21.4
)
