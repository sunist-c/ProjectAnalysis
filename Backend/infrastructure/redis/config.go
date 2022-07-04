package redis

import (
	"strconv"
	"time"
)

// Config the configuration structure of a redis server
type Config struct {
	Address       string `json:"address" yaml:"address"`
	Username      string `json:"username" yaml:"username"`
	Password      string `json:"password" yaml:"password"`
	DatabaseIndex string `json:"database_index" yaml:"database_index"`
	Timeout       string `json:"timeout" yaml:"timeout"`
	MaxRetry      string `json:"max_retry" yaml:"max_retry"`
}

// toDatabaseIndex exchange configuration field to database-index option
func (c Config) toDatabaseIndex() int {
	if index, err := strconv.Atoi(c.DatabaseIndex); err != nil {
		return 0
	} else {
		return index
	}
}

// toTimeout exchange configuration field to read/write/retry timeout options
func (c Config) toTimeout() time.Duration {
	if timeout, err := strconv.Atoi(c.Timeout); err != nil {
		return time.Second * 4
	} else {
		return time.Duration(timeout) * time.Second
	}
}

// toMaxRetry exchange configuration field to max_retry option
func (c Config) toMaxRetry() int {
	if retry, err := strconv.Atoi(c.MaxRetry); err != nil {
		return 3
	} else {
		return retry
	}
}
