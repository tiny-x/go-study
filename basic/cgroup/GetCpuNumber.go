package main

// Importing automaxprocs automatically adjusts GOMAXPROCS.
import (
	"fmt"
	_ "go.uber.org/automaxprocs/maxprocs"
	"runtime"
)

func main() {
	cpuCount := runtime.NumCPU()
	gomaxprocs := runtime.GOMAXPROCS(cpuCount)
	fmt.Println(cpuCount)
	fmt.Println(gomaxprocs)
}
