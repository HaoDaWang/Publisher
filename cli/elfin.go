package cli

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"publisher/utils"
	"strings"
)

const (
	TEST_ENV_1 = iota
	TEST_ENV_2
	LOCAL_TEST_ENV
)

var elfinCli = map[int]string{
	TEST_ENV_2: "testing2",
	TEST_ENV_1: "testing1",
	LOCAL_TEST_ENV: "local",
}

func Command(name string, arg ...string) {
	fmt.Println(append([]string{name}, arg...))
}

func Publish(branchName string, env int) []string {
	cmd := exec.Command(
		"elfin",
		"deploy",
		"--local",
		"-e",
		elfinCli[env],
		"-b",
		branchName,
	)

	var stderr bytes.Buffer

	cmd.Stderr = &stderr
	cmd.Dir = utils.GetProjectName()

	out, err := cmd.Output()

	//fmt.Println(utils.GetProjectName())

	if err != nil {
		log.Fatal(stderr.String(), err)
	}

	outs := strings.Split(string(out), "\n")

	return outs
}