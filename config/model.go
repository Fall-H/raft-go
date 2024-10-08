package config

type Config struct {
	Info      InfoConfig
	Heartbeat HeartbeatConfig
	Vote      VoteConfig
	Serve     ServeInfo
}

type InfoConfig struct {
	Type string
	Term int32
	Name string
}

type HeartbeatConfig struct {
	MasterIp string
	Port     string
	Serve    []ServeConfig
}

type VoteConfig struct {
	Port  string
	Serve []ServeConfig
}

type ServeConfig struct {
	Name string
	Ip   string
}

type ServeInfo struct {
	Ip string
}
