package types

import _ "gopkg.in/yaml.v3"

type Config struct {
	Server *ConfigServer `yaml:"server"`
	DB     *ConfigDb     `yaml:"db"`
}

type ConfigServer struct {
	Domain string `yaml:"domain"`
	Port   int    `yaml:"port"`
}

type ConfigDb struct {
	URL        string `yaml:"url"`
	Database   string `yaml:"database"`
	Collection string `yaml:"collection"`
}
