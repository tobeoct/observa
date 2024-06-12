package config

import (
	_ "embed"
	"fmt"

	"github.com/BurntSushi/toml"
)

type AlertConfiguration struct {
	Recipients []string `toml:"recipients"`
	Sender     string   `toml:"sender"`
}

type LogConfiguration struct {
	LogToConsole bool `toml:"logToConsole"`
}

type CollectorConfiguration struct {
	Url bool `toml:"url"`
}

type Config struct {
	Alert     AlertConfiguration     `toml:"alert"`
	Log       LogConfiguration       `toml:"log"`
	Collector CollectorConfiguration `toml:"collector"`
}

var Env Config

//go:embed env.toml
var envData []byte

func LoadConfig(path string, appMode string) error {
	var conf Config

	err := toml.Unmarshal(envData, &conf)
	if err != nil {
		fmt.Println(err.Error())
		return fmt.Errorf("error unmarshalling TOML: %v", err)
	}
	Env = conf
	return nil
}
