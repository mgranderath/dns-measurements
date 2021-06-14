package main

import (
	"bufio"
	"context"
	"github.com/mgranderath/dns-measurements/db"
	"github.com/mgranderath/dns-measurements/workflow"
	"golang.org/x/sync/semaphore"
	"log"
	"net"
	"os"
	"os/signal"
)

func main() {
	file, err := os.Open("./dns-list/out.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var sem = semaphore.NewWeighted(int64(10))
	defer db.Close()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func(){
		for range c {
			file.Close()
			db.Close()
			os.Exit(1)
		}
	}()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		ip := net.ParseIP(scanner.Text())
		if ip.To4() != nil {
			sem.Acquire(context.Background(), 1)
			go func() {
				log.Println("start for", ip.String())
				workflow.Standard(ip.To4().String())
				sem.Release(1)
			}()
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
