package main

import (
	"raft/config"
	"raft/heartbeat"
	"raft/model"
	"raft/observer"
	"raft/serve"
	"raft/vote"
	"sync"
	"time"
)

func init() {

	state := observer.NewItem(config.GConfig.Info.Type)

	heartbeatService := &heartbeat.Service{
		Client: model.Client{
			Lock:            sync.Mutex{},
			NextTimeOutTime: time.Now(),
			Ip:              config.GConfig.Heartbeat.Port,
		},
		ObserverState: state,
	}

	voteService := &vote.Service{
		ObserverState: state,
	}

	serveService := &serve.Service{
		ObserverState: state,
	}

	go heartbeatService.CreateRpc()
	go voteService.CreateVoteServe()
	go serveService.CreateServe()

	state.Register(heartbeatService)
	state.Register(voteService)
	state.Register(serveService)

	state.NotifyAll(config.GConfig.Info.Type)
}

func main() {
	for {

	}
}
