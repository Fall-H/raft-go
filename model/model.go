package model

import (
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
	GIpServe []IpServe
	//PollingTime  = 100 * time.Millisecond
	//Timeout      = 400 * time.Millisecond
	//IntervalTime = 200 * time.Millisecond
	PollingTime  = 10 * time.Second
	Timeout      = 40 * time.Second
	IntervalTime = 20 * time.Second
	MasterIp     = ""
)

const (
	StateVote   = "vote"
	StateVoted  = "voted"
	StateMaster = "master"
	StateSlave  = "slave"
)
