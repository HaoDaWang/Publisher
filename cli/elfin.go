package cli

import (
	"log"
	"os/exec"
	"strings"
)

const (
	TEST_ENV_1 = iota
	TEST_ENV_2
	LOCAL_TEST_ENV
)

func Publish(branchName string, env int) []string {
	cmd := exec.Command("yarn", "--help")
	out, err := cmd.Output()

	if err != nil {
		log.Fatal(err)
	}

	outs := strings.Split(string(out), "\n")

	return outs
}