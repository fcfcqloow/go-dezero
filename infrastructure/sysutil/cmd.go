package sysutil

import (
	"os/exec"
	"strings"
)

func run(cmdLine string) {
	args := strings.Split(cmdLine, " ")
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Run()
}
