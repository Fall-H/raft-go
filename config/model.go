package config

type Config struct {
	Info      InfoConfig
	Heartbeat HeartbeatConfig
	Vote      VoteConfig
}

type InfoConfig struct {
	Type string
	Term int32
	Name string
}

type HeartbeatConfig struct {
	Port  int
	Serve []ServeConfig
}

type VoteConfig struct {
	Port  int
	Serve []ServeConfig
}

type ServeConfig struct {
	Name string
	Ip   string
}
