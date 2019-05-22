package message

import (
	jsoniter "github.com/json-iterator/go"
	"os"
)

var property *Property

// Property ...
type Property struct {
	Welcome     string `json:"welcome"`
	GroupName   string `json:"group_name"`
	Host        string `json:"host"`
	HookAddress string `json:"hook_address"`
	Token       string `json:"token"`
}

// LoadProperty ...
func LoadProperty(pathname string) error {
	property = &Property{
		Welcome:     "欢迎 %s 加入%s, 新朋友请先看置顶消息, 老朋友快来围观这只野生的新人, 邀请更多好友加入求哈希, 大家一起加速.\n（本消息将于30秒后自动销毁）",
		GroupName:   "",
		Host:        "https://bot.dhash.app/",
		HookAddress: "crVuYHQbUWCerib3",
		Token:       "",
	}
	file, e := os.OpenFile(pathname, os.O_RDONLY, 0755)
	if e != nil {
		return e
	}
	dec := jsoniter.NewDecoder(file)
	e = dec.Decode(property)
	if e != nil {
		return e
	}
	return nil
}

// GetProperty ...
func GetProperty() *Property {
	return property
}
