
package main

import (
	"fmt"
	"os"
)

func main() {
	chapter2_3()
	chapter2_3_1_01()
	chapter2_3_1_02()
	chapter2_3_2()
}

func chapter2_3() {
	var myInt int
	var myBoolean bool
	var myString string
	var myMap map[string]int

	fmt.Printf("myInt : %d\n",myInt)
	fmt.Printf("myBoolean : %t\n",myBoolean)
	fmt.Printf("myString : %s\n",myString)
	fmt.Printf("myMap : %v\n",myMap)
}

func chapter2_3_1_01() {
	var b, f, s = true, 2.3, "four"
	fmt.Printf("b : %t\n",b)
	fmt.Printf("f : %f\n",f)
	fmt.Printf("s : %s\n",s)

	i, j := 111 , 222	// declaration
	i, j = j, i			// assignment
	fmt.Printf("i : %d\n",i)
	fmt.Printf("j : %d\n",j)
}

func chapter2_3_1_02() {
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

func chapter2_3_2() {
	// 01
	x := 1			// int x
	p := &x			// int* p
	*p = 2
	fmt.Println(x)	// 2
	fmt.Println(&x)	// &x
	fmt.Println(*p)	// 2
	fmt.Println(p)	// &x

	// 02
	var x1,y1 int
	fmt.Println(&x1 == &x1, &x1 == &y1, &x1 == nil)

	// 03
	fmt.Println(f())
	fmt.Println(f() == f())

	// 04
	v := 1
	incr(&v)				// 2
	fmt.Println(incr(&v))	// 3
}

func f() *int {
	v := 1
	return &v
}

func incr(p *int) int {
	*p++
	//(*p)++
	return *p
}