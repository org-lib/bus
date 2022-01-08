package main

import (
	"fmt"
	"github.com/org-lib/bus/xshell"
)

func main() {
	shell, err := xshell.Powershell()
	if err != nil {
		panic(err)
	}
	defer shell.Exit()

	// ... 和它交互
	stdout, stderr, err := shell.Execute("Get-WmiObject -Class Win32_Processor")
	if err != nil {
		panic(err)
	}

	fmt.Println(stdout)
	fmt.Println(stderr)
}
