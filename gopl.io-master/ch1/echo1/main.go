// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 4.
//!+

// Echo1 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)


func main() {
	//normalFunc()
	//forTest01()
	//forTest02()
	//example_1_1()
	//example_1_2()
	example_1_3()
}

func normalFunc() {
	var s, sep string
	var argLen = len(os.Args)
	for i := 1; i < argLen; i++ {
	//for i := 1; i < argLen; ++i {	// error
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}

func forTest01(){
	var s, sep string
	var argLen = len(os.Args)
	var i = 1
	for i < argLen {
		s += sep + os.Args[i]
		sep = " "
		i++
	}
	fmt.Println(s)
}

func forTest02(){
	var s, sep string
	var argLen = len(os.Args)
	var i = 1
	for {
		if !(i < argLen) {
			break
		}
		s += sep + os.Args[i]
		sep = " "
		i++
	}
	fmt.Println(s)
}

func example_1_1() {
	var s, sep string
	var argLen = len(os.Args)
	for i := 0; i < argLen; i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}

func example_1_2() {
	var s string
	var argLen = len(os.Args)
	for i := 1; i < argLen; i++ {
		//s += "arg[" + string(i) + "] : " + os.Args[i] + "\n"
		s += "arg[" + strconv.Itoa(i) + "] : " + os.Args[i] + "\n"
	}
	fmt.Println(s)
}

func example_1_3() {
	//TODO: test program (11.4)
	arrayTmpeString := []string{
		"start_string",
		"dddd1","dddd1","dddd1","dddd1","dddd1","dddd1","dddd1","dddd1","dddd1","dddd1",
		"dddd2","dddd2","dddd2","dddd2","dddd2","dddd2","dddd2","dddd2","dddd2","dddd2",
		"dddd3","dddd3","dddd3","dddd3","dddd3","dddd3","dddd3","dddd3","dddd3","dddd3",
		"dddd4","dddd4","dddd4","dddd4","dddd4","dddd4","dddd4","dddd4","dddd4","dddd4",
		"dddd5","dddd5","dddd5","dddd5","dddd5","dddd5","dddd5","dddd5","dddd5","dddd5",
		"dddd6","dddd6","dddd6","dddd6","dddd6","dddd6","dddd6","dddd6","dddd6","dddd6",
		"dddd7","dddd7","dddd7","dddd7","dddd7","dddd7","dddd7","dddd7","dddd7","dddd7",
		"end_string",
	}

	//append(arrayTmpeString, )



	// normal
	startTime := time.Now()
	var s, sep string
	var argLen = len(arrayTmpeString)
	for i := 0; i < argLen; i++ {
		s += sep + arrayTmpeString[i]
		sep = " "
	}
	fmt.Println(s)
	fmt.Printf("normal : %.2fs elapsed\n", time.Since(startTime).Seconds())

	fmt.Println("------------")


	// join
	startTime = time.Now()
	fmt.Println(strings.Join(arrayTmpeString[0:], " "))
	fmt.Printf("join : %.2fs elapsed\n", time.Since(startTime).Seconds())
}

func normalFuncForBench() {

}

func joinFuncForBench() {

}
//!-
