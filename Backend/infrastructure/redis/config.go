package redis

// Config the configuration structure of a redis server
type Config struct {
	Address       string `json:"address" yaml:"address"`
	Username      string `json:"username" yaml:"username"`
	Password      string `json:"password" yaml:"password"`
	DatabaseIndex string `json:"database_index" yaml:"database_index"`
}
