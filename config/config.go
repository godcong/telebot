package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	jsoniter "github.com/json-iterator/go"
)

type Config struct {
	Bot    Bot    `json:"bot"`
	Client Client `json:"client"`
	Debug  bool   `json:"debug"`
	Auth   string `json:"auth"`
}

// LoadConfig ...
func LoadConfig(pathname string) (*Config, error) {
	property := defaultConfig()
	file, err := os.Open(pathname)
	if err != nil {
		err = writeDefaultConfig(pathname)
		return property, err
	}
	dec := jsoniter.NewDecoder(file)
	err = dec.Decode(property)
	if err != nil {
		return nil, err
	}
	log.Printf("property:%+v", *property)
	return property, nil
}

func writeDefaultConfig(pathname string) error {
	data, err := json.MarshalIndent(defaultConfig(), "", " ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(pathname, data, 0755)
}

func defaultConfig() *Config {
	return &Config{
		Bot: Bot{
			Model:          BotModelWebhook,
			Welcome:        "",
			GroupName:      "",
			BotName:        "",
			Host:           "",
			HookAddress:    "",
			Token:          "",
			Download:       "",
			Recognition:    "",
			Database:       "",
			Point:          0,
			RecognitionCMD: "",
			KnownPath:      "",
			Rule:           "",
			LocalURL:       "",
		},
		Client: Client{
			APIID:   0,
			APIHash: "",
		},
	}
}
