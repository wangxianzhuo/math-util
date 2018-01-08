package config

import (
	"encoding/json"
	"io/ioutil"
)

func Load(config interface{}, file string) error {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, config)
}
