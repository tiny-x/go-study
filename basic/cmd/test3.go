package cmd

import (
	"fmt"
	"os"
	"os/exec"
)

func test_xxxx() {

	//command := `java -jar a.jar`
	//cmd := exec.Command("/bin/sh", "-c", command)

	cmd := exec.Command("C:/Java/xx xx/bin/java", "-jar", "a.jar")

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
