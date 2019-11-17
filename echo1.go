package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	//  echo 1
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)

	//echo 2
	var s1, sep1 string

	for _, v := range os.Args[1:] {
		s1 += sep1 + v
		sep1 = " "
	}
	fmt.Println(s1)
	//	echo 3
	fmt.Println(strings.Join(os.Args[1:], " "))
}

// i++  i-- 不是一个表达式

// for  init , condition, post post语句在循环地之后执行
// for condition  传统的while循环
// for condition  传统的while循环
// for  无限循环
// for  _, v range := os.Args[1:]
//  strings.Join(os.Args[1:]," ")
