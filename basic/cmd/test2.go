package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {

	command := `[ -e /home/chaos_filldisk.log.dat ] && echo true || echo false`
	command = `ls /opt`
	//command = strings.Replace(command, " ", `\ `, -1)
	cmd := exec.Command("/bin/sh", "-c", command)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	environ := os.Environ()
	fmt.Println(environ)
	cmd.CombinedOutput()
	fmt.Println(cmd.Env)

	if err := cmd.Start(); err != nil {
		fmt.Println(err)
	}

	if err := cmd.Wait(); err != nil {
		fmt.Println(err)
	}

}
