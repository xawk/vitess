package main

import "github.com/xawk/vitess/go/stats/statsd"

func init() {
	statsd.Init("vttablet")
}
