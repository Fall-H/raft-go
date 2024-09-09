package heartbeat

import (
	config "raft/config"
	"raft/model"
)

func init() {
	for _, serve := range config.GConfig.Heartbeat.Serve {
		GIpServe = append(GIpServe, model.IpServe{
			Ip:      serve.Ip,
			Name:    serve.Name,
			IsAlive: true,
		})
	}
}
