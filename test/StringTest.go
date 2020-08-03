package main

import (
	"fmt"

	"regexp"

	"os"
	"strconv"

)

func main()  {


	//目标字符串
	searchIn := `
./blade create file append --filepath ~/chaosblade/mock.log 
--content "\@{date '+%Y-%m-%d %H:%M\:%S'\}} @{date '+%Y-%m-%d %H:%M:%S' @{\n} {@{x\\\x} @{date '+%Y-%m-%d %H:%M:%Sxx'}hell world" 
--count 2 --interval=2 @[RANDOM * 10] \@JAVA_HOME`

	//解释正则表达式
	reg := regexp.MustCompile(`\\?@\{(?s:([^(@{})]*[^\\]))\}|\\?@\[((?s:[^(@{})]*[^\\]))\]|\\?@\w+`)

	//提取关键信息
	result := reg.FindAllStringSubmatch(searchIn, -1)

	//过滤<></>
	for _, text := range result {
		fmt.Println("text[o] = ", text[0])
		fmt.Println("text[1] = ", text[1])
		if strings.HasPrefix(text[0], "\\@") {
			searchIn = strings.Replace(searchIn, text[0], text[0][1 : len(text[0]) - 1], 1)
			continue
		}


	}

	fmt.Println(searchIn)

	fileInfo, _ := os.Stat("/temp/1.log")
	formatInt := strconv.FormatInt(int64(fileInfo.Mode().Perm()), 8)
	fmt.Print(formatInt)
	fmt.Print(formatInt)
	fmt.Print(formatInt)


}
