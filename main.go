package main

import (
	"flag"
	"io/ioutil"
	"log"
)

func main() {
	flag.Parse()

	token, e := ioutil.ReadFile("token")
	if e != nil {
		return
	}
	log.Println(string(token))
	BootWithGAE(string(token))
}
