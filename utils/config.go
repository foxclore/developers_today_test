package utils

import (
	"encoding/json"
	"os"
)

func ReadConfig(parseTo interface{}, names ...string) error {
	name := UnpackDefault(names, "config.json")

	d, err := os.ReadFile(name)

	if err != nil {
		return err
	}

	return json.Unmarshal(d, parseTo)
}
