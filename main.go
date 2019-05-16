package main

import (
	"github.com/girlvr/yinhe_bot/message"
	log "github.com/godcong/go-trait"
	_ "github.com/mattn/go-sqlite3"
	"io/ioutil"
)

func main() {
	token, e := ioutil.ReadFile("token")
	if e != nil {
		return
	}
	log.InitGlobalZapSugar()
	message.BootWithGAE(string(token))
}
