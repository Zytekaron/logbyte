package config

import (
	"errors"
	"gopkg.in/yaml.v3"
	"logbyte/src/types"
	"os"
)

var configDirs = []string{
	"/etc/opt/endpoints.yml",
	"config.yml",
}

func Load() (cfg *types.Config, err error) {
	var file *os.File
	for i := 0; i < len(configDirs); i++ {
		file, err = os.Open(configDirs[i])
		if err != nil {
			continue
		}

		err = yaml.NewDecoder(file).Decode(&cfg)
		return
	}

	return nil, errors.New("could not find any config files")
}
