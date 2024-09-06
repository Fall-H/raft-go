package heartbeat

import (
	config "raft/config"
)

func init() {
	for _, serve := range config.GConfig.Serve {
		GIpServe = append(GIpServe, IpServe{
			Ip:      serve.Ip,
			Name:    serve.Name,
			IsAlive: true,
		})
	}
}
