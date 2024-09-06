package vote

import "raft/config"

func init() {
	for _, serve := range config.GConfig.Heartbeat.Serve {
		GIpServe = append(GIpServe, IpServe{
			Ip:      serve.Ip,
			Name:    serve.Name,
			IsAlive: true,
		})
	}
}
