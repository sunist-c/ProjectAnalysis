package oauth

// MiddlewareConfig the structure of oauth-application config
type MiddlewareConfig struct {
	Database     DatabaseConfig `json:"database" yaml:"database"`
	Redis        RedisConfig    `json:"redis" yaml:"redis"`
	EnableServer bool           `json:"enable_server" yaml:"enable_server"`
	ServerRoute  string         `json:"server_route" yaml:"server_route"`
}

// DatabaseConfig the structure of oauth-database config
type DatabaseConfig struct {
	Address  string `json:"address" yaml:"address"`
	Username string `json:"username" yaml:"username"`
	Password string `json:"password" yaml:"password"`
	Database string `json:"database" yaml:"database"`
}

// RedisConfig the structure of oauth-cache config
type RedisConfig struct {
	Address  string `json:"address" yaml:"address"`
	Username string `json:"username" yaml:"username"`
	Password string `json:"password" yaml:"password"`
	Database string `json:"database" yaml:"database"`
}
