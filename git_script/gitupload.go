package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"time"
)

func execute() {

	outStatus, err := exec.Command("powershell", "git status").Output()
	outputStatus := string(outStatus[:])
	fmt.Println(outputStatus)

	time.Sleep(5 * time.Second)

	outAdd, err := exec.Command("powershell", "git add .").Output()
	outputAdd := string(outAdd[:])
	fmt.Println(outputAdd)

	time.Sleep(5 * time.Second)

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter commit message: ")
	commitMessage, _ := reader.ReadString('\n')

	outCommit, err := exec.Command("powershell", fmt.Sprintf("git commit -am \"%s\"", commitMessage)).Output()
	outputCommit := string(outCommit[:])
	fmt.Println(outputCommit)

	time.Sleep(5 * time.Second)

	outPush, err := exec.Command("powershell", "git push").Output()
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
