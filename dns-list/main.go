package main

import (
	"bufio"
	"github.com/mgranderath/dns-measurements/workflow"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	file, err := os.Open("./dns-list/nameservers.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		ip := net.ParseIP(scanner.Text())
		if ip.To4() != nil {
			workflow.StartForIP(ip.To4().String())
			time.Sleep(time.Second)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
