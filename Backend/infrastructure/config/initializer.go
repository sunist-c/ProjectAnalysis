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
		buffer, err := ioutil.ReadFile(filePath)
		if err != nil {
			return
		} else {
			return json.Unmarshal(buffer, receiver)
		}
	case ".yml":
		buffer, err := ioutil.ReadFile(filePath)
		if err != nil {
			return
		} else {
			return yaml.Unmarshal(buffer, receiver)
		}
	case ".yaml":
		buffer, err := ioutil.ReadFile(filePath)
		if err != nil {
			return
		} else {
			return json.Unmarshal(buffer, receiver)
		}
	case ".ini":
		file, err := ini.Load(filePath)
		if err != nil {
			return
		} else {
			return file.MapTo(receiver)
		}
	default:
		return errors.New("unknown config file extension")
	}
}
