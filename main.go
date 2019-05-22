package main

import (
	"flag"
	"github.com/girlvr/yinhe_bot/message"
	_ "github.com/mattn/go-sqlite3"
)

var path = flag.String("path", "yinhe.json", "default property path")

func main() {
	flag.Parse()
	message.BootWithGAE(*path)
}
