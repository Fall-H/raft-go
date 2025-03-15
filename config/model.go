package config

type Config struct {
	Info        InfoConfig
	Heartbeat   HeartbeatConfig
	Vote        VoteConfig
	Serve       ServeInfo
	MasterServe MasterServeConfig
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

type MasterServeConfig struct {
	Ip string
}

type VoteConfig struct {
	Port  string
	Serve []ServeConfig
}

type ServeConfig struct {
	Ip   string
	Name string
}

type ServeInfo struct {
	Ip string
}
