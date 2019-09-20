package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

func execute() {

	// out1, err := exec.Command("powershell", "git status").Output()
	// output1 := string(out1[:])
	// fmt.Println(output1)

	// time.Sleep(5 * time.Second)

	// out2, err := exec.Command("powershell", "git add .").Output()
	// output2 := string(out2[:])
	// fmt.Println(output2)

	// time.Sleep(5 * time.Second)

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter commit description: ")
	text, _ := reader.ReadString('\n')

	out3, err := exec.Command("powershell", fmt.Sprintf("git commit -am \"%s\"", text)).Output()
	output3 := string(out3[:])
	fmt.Println(output3)

	if err != nil {
		fmt.Printf("%s", err)
	} else {
		fmt.Println("Command Successfully Executed")
	}

}

func main() {
	execute()
}
