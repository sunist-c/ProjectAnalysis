package mysql

// Config the configuration structure of a mysql database
type Config struct {
	Address           string `json:"address" yaml:"address"`
	Username          string `json:"username" yaml:"username"`
	Password          string `json:"password" yaml:"password"`
	DatabaseName      string `json:"database_name" yaml:"database_name"`
	MaxOpenConnection string `json:"max_open_connection" yaml:"max_open_connection"`
	MaxIdleConnection string `json:"max_idle_connection" yaml:"max_idle_connection"`
}
