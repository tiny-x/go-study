package net

import (
	"fmt"
	"net"
	"testing"
)

func TestA(t *testing.T) {

	interfaces, err := net.Interfaces()
	if err != nil {
		panic(interfaces)
	}

	for _, i := range interfaces {
		addrs, err := i.Addrs()
		if err != nil {
			panic(err)
		}
		for _, addr := range addrs {
			if a, ok := addr.(*net.IPNet); ok {
				fmt.Print(i.Name, a.IP.String(), a.IP.To4().IsLoopback(), "\t\n")
			}
		}
	}
}
