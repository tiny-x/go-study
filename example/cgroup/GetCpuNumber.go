package main

// Importing automaxprocs automatically adjusts GOMAXPROCS.
import (
	"fmt"
	_ "go.uber.org/automaxprocs/maxprocs"
	"runtime"
)

// To render a whole-file example, we need a package-level declaration.
var _ = ""

func main() {
	size := runtime.GOMAXPROCS(0)
	fmt.Print(size)
}
