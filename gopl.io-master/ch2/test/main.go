
package main

import (
	"fmt"
	"os"
)

func main() {
	test01()
	test02()
	test03()
}

func test01() {
	var myInt int
	var myBoolean bool
	var myString string
	var myMap map[string]int

	fmt.Printf("myInt : %d\n",myInt)
	fmt.Printf("myBoolean : %t\n",myBoolean)
	fmt.Printf("myString : %s\n",myString)
	fmt.Printf("myMap : %v\n",myMap)
}

func test02() {
	var b, f, s = true, 2.3, "four"
	fmt.Printf("b : %t\n",b)
	fmt.Printf("f : %f\n",f)
	fmt.Printf("s : %s\n",s)

	i, j := 111 , 222	// declaration
	i, j = j, i			// assignment
	fmt.Printf("i : %d\n",i)
	fmt.Printf("j : %d\n",j)
}

func test03() {
	// 01
	name := "tmp.txt"
	f, err := os.Open(name)
	if err != nil {
		return
	}
	f.Close()

	// 02
	inFileName := "tmp2.txt"
	outFileName := "tmp3.txt"
	in, err := os.Open(inFileName)
	if err != nil {return}
	in.Close()
	out, err := os.Create(outFileName)
	if err != nil {return}
	out.Close()

	// 03
	f2, err := os.Open(inFileName)
	f2.Close()


	// 03-1 : compile error - not any declaration
	//f2, err := os.Create(outFileName)

	// 03-2 : err2 is declaration
	//f2, err2 := os.Create(outFileName)
	//if err2 != nil {return}

	// 03-3 : all assignment
	f2, err = os.Create(outFileName)

	f2.Close()
}
