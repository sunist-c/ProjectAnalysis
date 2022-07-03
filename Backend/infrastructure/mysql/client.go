package mysql

import (
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

// Client the client to operate mysql database
type Client struct {
	connection *xorm.Engine
	connected  bool
}

// Connect try to connect mysql database with configuration
func (c *Client) Connect(cfg Config) (err error) {
	c.connection, err = xorm.NewEngine("mysql", fmt.Sprintf("%v:%v@tcp(%v)/%v?charset=utf8mb4", cfg.Username, cfg.Password, cfg.Address, cfg.DatabaseName))
	if err != nil {
		return
	}

	c.connection.SetMaxOpenConns(cfg.toMaxOpenConn())
	c.connection.SetMaxIdleConns(cfg.toMaxIdleConn())

	err = c.connection.Ping()
	return
}

// Sync try to sync table structures between code and database
func (c Client) Sync(opt interface{}) (err error) {
	return c.connection.Sync2(opt)
}

// Create insert records into mysql database
func (c Client) Create(opts ...interface{}) (effected []int, err error) {
	if !c.connected {
		return nil, errors.New("database does not connected")
	}

	effected = make([]int, 0, 8)
	for i, opt := range opts {
		_, err = c.connection.InsertOne(opt)
		if err != nil {
			return effected, err
		} else {
			effected = append(effected, i)
		}
	}
	return effected, nil
}

// Update update a record into mysql database
func (c Client) Update(newOpt, oldOpt interface{}) (success bool, err error) {
	if !c.connected {
		return false, errors.New("database does not connected")
	}

	_, err = c.connection.Update(newOpt, oldOpt)
	if err != nil {
		return false, err
	} else {
		return true, nil
	}
}

// Retrieve query a record with conditions in mysql database
func (c Client) Retrieve(condition interface{}) (success bool, err error) {
	if !c.connected {
		return false, errors.New("database does not connected")
	}

	return c.connection.Get(&condition)
}

// Delete delete a record with conditions in mysql database
func (c Client) Delete(condition interface{}) (success bool, err error) {
	if !c.connected {
		return false, errors.New("database does not connected")
	}

	_, err = c.connection.Delete(condition)
	if err != nil {
		return false, err
	} else {
		return true, nil
	}
}
