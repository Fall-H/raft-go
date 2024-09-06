package heartbeat

import (
	"sync"
	"time"
)

type Client struct {
	lock            sync.Mutex
	nextTimeOutTime time.Time
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
)

const timeout = 10
const intervalTime = 9
