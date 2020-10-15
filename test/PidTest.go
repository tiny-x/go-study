package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/chaosblade-io/chaosblade-spec-go/util"
)

func main() {
	name := "/Users"
	exist := util.IsDir(name)
	fmt.Println(exist)

}

func md5Hex(s string) string {
	m := md5.New()
	m.Write([]byte(s))
	return hex.EncodeToString(m.Sum(nil))
}
