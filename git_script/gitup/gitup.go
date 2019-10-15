package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"time"
)

func execute() {

	var commitMessage string
	var branchName string

	flag.StringVar(&commitMessage, "m", "", "Input commit message (Required)")
	flag.StringVar(&branchName, "b", "", "Input branch name - master or other (Required)")
	flag.Parse()

	if commitMessage == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	if branchName == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	outStatus, err := exec.Command("powershell", "git status").Output()
	outputStatus := string(outStatus[:])
	fmt.Println(outputStatus)

	time.Sleep(3 * time.Second)

	outAdd, err := exec.Command("powershell", "git add .").Output()
	outputAdd := string(outAdd[:])
	fmt.Println(outputAdd)

	time.Sleep(5 * time.Second)

	outCommit, err := exec.Command("powershell", fmt.Sprintf("git commit -am \"%s\"", commitMessage)).Output()
	outputCommit := string(outCommit[:])
	fmt.Println(outputCommit)

	time.Sleep(5 * time.Second)

	outPush, err := exec.Command("powershell", fmt.Sprintf("git push origin \"%s\"", branchName)).Output()
	outputPush := string(outPush[:])
	fmt.Println(outputPush)

	if err != nil {
		fmt.Printf("%s", err)
	} else {
		fmt.Println("Command Successfully Executed")
	}

}

func main() {
	execute()
}
