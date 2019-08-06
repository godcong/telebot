package main

import (
	"flag"

	"github.com/glvd/bot/message"
	_ "github.com/mattn/go-sqlite3"
)

var path = flag.String("path", "yinhe.json", "default property path")
var port = flag.String("port", "443", "default port")

func main() {
	flag.Parse()
	message.BootWithGAE(*path, *port)
}
