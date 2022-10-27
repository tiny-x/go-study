package protocol

import (
	"fmt"
	"testing"
)

func TestStart(t *testing.T) {
	err := Start(20000)
	if err != nil {
		fmt.Println(err)
	}
}
