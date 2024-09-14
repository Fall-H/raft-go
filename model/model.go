package model

import (
	"raft/util"
	"sync"
	"time"
)

type Client struct {
	Lock            sync.Mutex
	NextTimeOutTime time.Time
	Ip              string
}

type Pong struct {
	Ip      string
	message string
}

type IpServe struct {
	Ip      string
	Name    string
	IsAlive bool
}

var (
	//PollingTime  = 100 * time.Millisecond
	//Timeout      = 400 * time.Millisecond
	//IntervalTime = 200 * time.Millisecond
	PollingTime = 5 * time.Second
	//Timeout      = util.RandomTime(150, 300) * time.Second
	Timeout      = util.RandomTime(10, 20)
	IntervalTime = 5 * time.Second
	Number       = int32(0)
	OpList       []string
)

const (
	StateVote   = "vote"
	StateVoted  = "voted"
	StateMaster = "master"
	StateSlave  = "slave"
)
