package message

import (
	"os"

	jsoniter "github.com/json-iterator/go"
)

var property *Property

// Property ...
type Property struct {
	Welcome        string  `json:"welcome"`
	GroupName      string  `json:"group_name"`
	BotName        string  `json:"bot_name"`
	Host           string  `json:"host"`
	HookAddress    string  `json:"hook_address"`
	Token          string  `json:"token"`
	Download       string  `json:"download"`
	Recognition    string  `json:"recognition"`
	Database       string  `json:"database"`
	Point          float64 `json:"point"`
	RecognitionCMD string  `json:"recognition_cmd"`
	KnownPath      string  `json:"known_path"`
	Rule           string  `json:"rule"`
	LocalURL       string  `json:"local_url"`
}

// LoadProperty ...
func LoadProperty(pathname string) error {
	property = &Property{}
	file, e := os.Open(pathname)
	if e != nil {
		return e
	}
	dec := jsoniter.NewDecoder(file)
	e = dec.Decode(property)
	if e != nil {
		return e
	}
	log.Infof("property:%+v", *property)
	return nil
}

// GetProperty ...
func GetProperty() *Property {
	return property
}
