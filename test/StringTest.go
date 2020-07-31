package main

import (
	"fmt"
	"os"
	"strconv"
)

func main()  {

	fileInfo, _ := os.Stat("/temp/1.log")
	formatInt := strconv.FormatInt(int64(fileInfo.Mode().Perm()), 8)
	fmt.Print(formatInt)
	fmt.Print(formatInt)
	fmt.Print(formatInt)

}
