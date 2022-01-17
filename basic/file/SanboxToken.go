package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"regexp"
	"strings"
)

func main() {

	osCmd := fmt.Sprintf("grep %s", fmt.Sprintf(`%s %s | grep %s | tail -1 | awk -F ";" '{print $3";"$4}'`,
		"27599253248", getSandboxTokenFile(), "defalut"))

	fmt.Println(osCmd)

	bytes, _ := ioutil.ReadFile(getSandboxTokenFile())
	findString := regexp.MustCompile(`(default;384344107648;localhost;)[0-9]+`).FindString(string(bytes))
	split := strings.Split(strings.TrimSpace(string(bytes)), "\n")
	fmt.Println(findString)

	row := split[len(split)-1]
	rows := strings.Split(row, ";")
	fmt.Println(rows[len(rows)-1])
}

func getSandboxTokenFile() string {
	userHome, _ := os.UserHomeDir()
	return path.Join(userHome, ".sandbox.token")
}
