package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/godcong/go-trait"
)

var log = trait.NewZapSugar()

func main() {
	args := os.Args
	dir, err := os.Getwd()
	if len(args) > 1 {
		err = nil
		dir = args[1]
	}
	if err != nil {
		log.Info("wd:", err)
		return
	}
	fp, e := filepath.Abs(dir)
	if e != nil {
		log.Error(e)
		return
	}
	a := fmt.Sprintf("--tolerance 0.54 /home/ubuntu/face_pic/ %s | cut -d ',' -f2", fp)

	cmd := exec.Command("face_recognition", strings.Split(a, " ")...)

	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	log.Info("combined out:\n%s\n", string(out))

}
