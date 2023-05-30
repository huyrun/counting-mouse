package config

import (
	"encoding/json"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/v2"
)

type AppConfig struct {
	ServerAddress string `json:"server_address"`
}

func NewAppConfig() (*AppConfig, error) {
	var k = koanf.New("")
	if err := k.Load(env.Provider("", "", nil), nil); err != nil {
		return nil, err
	}

	configMap := k.All()

	configJsonData, err := json.Marshal(configMap)
	if err != nil {
		return nil, err
	}

	var config AppConfig
	err = json.Unmarshal(configJsonData, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
