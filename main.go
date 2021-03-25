package main

import "github.com/mgranderath/dns-measurements/workflow"

func main() {
	workflow.StartForIP("8.8.8.8")
	workflow.StartForIP("94.140.14.14")
}