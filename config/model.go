package config

type Config struct {
	Info      InfoConfig
	Heartbeat HeartbeatConfig
	Serve     []ServeConfig
}

type InfoConfig struct {
	Type string
}

type HeartbeatConfig struct {
	Port int
}

type ServeConfig struct {
	Name string
	Ip   string
}
