package config

import (
	"encoding/json"
	"errors"
	"gopkg.in/ini.v1"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path/filepath"
)

type Configuration struct {
}

// Load read config file and map to provided structure
func (c Configuration) Load(filePath string, receiver interface{}) (err error) {
	switch filepath.Ext(filePath) {
	case ".json":
		buffer, err1 := ioutil.ReadFile(filePath)
		if err1 != nil {
			return err1
		} else {
			return json.Unmarshal(buffer, receiver)
		}
	case ".yml":
		buffer, err2 := ioutil.ReadFile(filePath)
		if err2 != nil {
			return err2
		} else {
			return yaml.Unmarshal(buffer, receiver)
		}
	case ".yaml":
		buffer, err3 := ioutil.ReadFile(filePath)
		if err3 != nil {
			return err3
		} else {
			return json.Unmarshal(buffer, receiver)
		}
	case ".ini":
		file, err4 := ini.Load(filePath)
		if err4 != nil {
			return err4
		} else {
			return file.MapTo(receiver)
		}
	default:
		return errors.New("unknown config file extension")
	}
}
