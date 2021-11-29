package config

import (
	"log"
	"os"

	jsoniter "github.com/json-iterator/go"
)

type Config struct {
	Bot    Bot    `json:"bot"`
	Client Client `json:"client"`
	Debug  bool   `json:"debug"`
}

// LoadConfig ...
func LoadConfig(pathname string) (*Config, error) {
	property := defaultConfig()
	file, e := os.Open(pathname)
	if e != nil {
		return nil, e
	}
	dec := jsoniter.NewDecoder(file)
	e = dec.Decode(property)
	if e != nil {
		return nil, e
	}
	log.Printf("property:%+v", *property)
	return property, nil
}

func defaultConfig() *Config {
	return &Config{
		Bot: Bot{
			Model:          BotModelWebhook.String(),
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
			APIID:   "",
			APIHash: "",
		},
	}
}
