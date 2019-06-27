package main

import (
	"flag"
	_ "github.com/mattn/go-sqlite3"
	"github.com/yinhevr/yinhe_bot/message"
)

var path = flag.String("path", "yinhe.json", "default property path")

func main() {
	flag.Parse()
	message.BootWithGAE(*path)
}
