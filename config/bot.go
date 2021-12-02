package config

//BotModel
//ENUM(webhook,update,max)
type BotModel int

// Bot ...
type Bot struct {
	Model          BotModel `json:"model"`
	Welcome        string   `json:"welcome"`
	GroupName      string   `json:"group_name"`
	BotName        string   `json:"bot_name"`
	Host           string   `json:"host"`
	HookAddress    string   `json:"hook_address"`
	Token          string   `json:"token"`
	Download       string   `json:"download"`
	Recognition    string   `json:"recognition"`
	Database       string   `json:"database"`
	Point          float64  `json:"point"`
	RecognitionCMD string   `json:"recognition_cmd"`
	KnownPath      string   `json:"known_path"`
	Rule           string   `json:"rule"`
	LocalURL       string   `json:"local_url"`
}
