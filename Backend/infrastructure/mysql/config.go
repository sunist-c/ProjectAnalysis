package mysql

import "strconv"

// Config the configuration structure of a mysql database
type Config struct {
	Address           string `json:"address" yaml:"address"`
	Username          string `json:"username" yaml:"username"`
	Password          string `json:"password" yaml:"password"`
	DatabaseName      string `json:"database_name" yaml:"database_name"`
	MaxOpenConnection string `json:"max_open_connection" yaml:"max_open_connection"`
	MaxIdleConnection string `json:"max_idle_connection" yaml:"max_idle_connection"`
}

// toMaxOpenConn exchange config fields to MAX_OPEN_CONNECTIONS option
func (c Config) toMaxOpenConn() int {
	if maxOpenConn, err := strconv.Atoi(c.MaxOpenConnection); err != nil {
		return 16
	} else {
		return maxOpenConn
	}
}

// toMaxIdleConn exchange config fields to MAX_IDLE_CONNECTIONS option
func (c Config) toMaxIdleConn() int {
	if maxIdleConn, err := strconv.Atoi(c.MaxIdleConnection); err != nil {
		return 16
	} else {
		return maxIdleConn
	}
}
