package parser

import (
	"os"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Directories []string `toml:"directories"`
}

func ParseConfig() (Config, error) {

	var config Config

	// read the TOML file
	file, err := os.ReadFile("config.toml")
	if err != nil {
		return config, err
	}

	// decode config contents
	if _, err := toml.Decode(string(file), &config); err != nil {
		return config, err
	}

	return config, nil
}
