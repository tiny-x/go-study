package disk

import (
	"fmt"
	"github.com/ncw/directio"
	"os"
	"os/exec"
	"testing"
)

func Test_disk_full(t *testing.T) {
	// 10M
	var block = 1024 * 1024 * 10
	blocks := make([]byte, block)
	f, err := directio.OpenFile("/tmp/1.log", os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Print(err)
	}
	for i := 0; i < 1024; i++ {
		f.Write(blocks)
	}
	f.Close()
}

func Test_disk_full_with_dd(t *testing.T) {
	// 10M
	command := exec.Command("dd", "if=/dev/zero", "of=/tmp/1.log",
		"bs=10M", "count=1024", "iflag=fullblock")
	output, err := command.CombinedOutput()
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println(string(output))
}

func Test_disk_full2(t *testing.T) {
	f, err := os.OpenFile("/tmp/1.log", os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Print(err)
	}
	err = Fallocate(f, 0, 1024*1024*100)
	if err != nil {
		fmt.Print(err)
	}
}
