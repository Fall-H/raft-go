package config

type Config struct {
	Heartbeat HeartbeatConfig `json:"heartbeat"`
}

type HeartbeatConfig struct {
	Port int `json:"port"`
}
