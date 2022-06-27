package commands

import (
	"bufio"
	"fmt"
	"os/exec"
)

// ReturnShellExecute 执行 shell 命令,（有的命令需要返回输出自定义结果值）
func ReturnShellExecute(command string) (string, error) {
	cmd := exec.Command("/bin/sh", "-c", command)
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(output), nil
}

// Execute 调用多参数二进制程序，或者脚本。如pt-grants,mydumper,等
func Execute(command string, param []string) error {
	cmd := exec.Command(command)
	cmd.Args = param

	cmdReader, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}
	scanner := bufio.NewScanner(cmdReader)
	go func(scanner *bufio.Scanner) {
		for scanner.Scan() {
			content := scanner.Text()
			fmt.Printf("%v\n", content)
		}
	}(scanner)

	err = cmd.Start()
	if err != nil {
		return err
	}

	err = cmd.Wait()
	if err != nil {
		return err
	}
	return nil
}
