package main

import (
	"github.com/sirupsen/logrus"
	"io/ioutil"
)

func main() {
	token, e := ioutil.ReadFile("token")
	if e != nil {
		return
	}
	logrus.Info(string(token))
	logrus.SetReportCaller(true)
	BootWithGAE(string(token))
}
