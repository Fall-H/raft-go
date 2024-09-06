package main

import "raft/heartbeat"

func init() {
	go heartbeat.StartPingService()
	go heartbeat.StartRpc()
}

func main() {
	for {

	}
}
