package main

import (
	"flag"
	"io/ioutil"
)

func main() {
	flag.Parse()

	token, e := ioutil.ReadFile("token")
	if e != nil {
		return
	}

	BootWithGAE(string(token))
}
