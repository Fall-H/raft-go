package vote

import (
	"raft/config"
	"raft/model"
)

func init() {
	for _, serve := range config.GConfig.Vote.Serve {
		GIpServe = append(GIpServe, model.IpServe{
			Ip:      serve.Ip,
			Name:    serve.Name,
			IsAlive: true,
		})
	}
}
