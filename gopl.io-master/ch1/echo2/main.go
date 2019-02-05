// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 6.
//!+

// Echo2 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	//normalMain()
	test01()
}

func normalMain() {
	//strTmp := ""				// only function able
	//var strTmp string
	//var strTmp = ""
	//var strTmp string = ""

	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

func test01() {
	for index, arg := range os.Args[1:] {
		fmt.Println("arg[" + strconv.Itoa(index) + "] : " + arg)
	}

	//for index, arg := range os.Args[1:] {
	//	fmt.Println("arg[" + strconv.Itoa(index) + "] : " + arg)
	//}
}
//!-
