

// https://golang.org/pkg/fmt/#hdr-Printing

// https://golang.org/ref/spec#Keywords
// https://golang.org/ref/spec#Constants
// https://golang.org/ref/spec#Types
// https://golang.org/ref/spec#Built-in_functions

package main

import (
	"fmt"
	"os"
	"./tempconv"
	"./tmppkg"
	//"tmppkg"
	//"gopl.io/ch2/tempconv"
)


func main() {
	chapter2_3()
	chapter2_3_1_01()
	chapter2_3_1_02()
	chapter2_3_2()
	chapter2_3_3()
	chapter2_4_1()
	chapter2_5()
	chapter2_6()
	chapter2_6_2()
}

func chapter2_3() {
	fmt.Println("\nchapter2_3 function call")
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
	fmt.Println("\nchapter2_3_1_01 function call")
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
	fmt.Println("\nchapter2_3_1_02 function call")
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
	fmt.Println("chapter2_3_2 function call")
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

func chapter2_3_3() {
	fmt.Println("\nchapter2_3_3 function call")

	// 01
	p := new(int)	// *int
	fmt.Println(*p)	// 0
	*p = 2
	fmt.Println(*p)	// 2

	// 02
	p2 := new(int)
	q2 := new(int)
	fmt.Println(p2 == q2)

	p3 := newInt01()
	q3 := newInt02()
	fmt.Println(p3)
	fmt.Println(*p3)
	fmt.Println(q3)
	fmt.Println(*q3)

	// 03
	fmt.Println(delta(10,20))
}

func newInt01() *int {
	return new(int)
}

func newInt02() *int{
	var dummy int
	return &dummy
}

func delta(old, new int) int {
	//testValue := new(int)	// error - already redefined 'new'
	return new - old
}

var global *int
func f2() {
	var x int	// 'x' heap
	x = 1
	global = &x
}

func g2() {
	y := new(int)	// 'y' stack
	*y = 1
}

func chapter2_4_1() {
	fmt.Println("\nchapter2_4_1 function call")
	fmt.Printf("gcd - %d\n",gcd(20,3))
	fmt.Printf("fib - %d\n",fib(6))
}

func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

func fib(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x + y
	}
	return x
}

func chapter2_5() {
	fmt.Println("\nchapter2_5 function call")
	Example_one()
	Example_two()

	var myTmp tmppkg.MyType = 10
	//fmt.Println(myTmp.myFunc())
	fmt.Println(myTmp.MyFunc())
	fmt.Println(myTmp.String())

}

func Example_one() {
	fmt.Printf("%g\n", tempconv.BoilingC-tempconv.FreezingC)			// "100"
	fmt.Printf("%v\n", tempconv.BoilingC-tempconv.FreezingC)			// "100°C"
	boilingF := tempconv.CToF(tempconv.BoilingC)
	fmt.Printf("%g\n", boilingF-tempconv.CToF(tempconv.FreezingC))	// "180"
	fmt.Println(boilingF-tempconv.CToF(tempconv.FreezingC))					// "180°F"
	//fmt.Printf("%g\n", boilingF-tempconv.FreezingC)						// compile error: type mismatch
}

func Example_two() {
	c := tempconv.FToC(212.0)
	fmt.Println(c.String())			// "100°C"
	fmt.Printf("%v\n", c)	// "100°C"; no need to call String explicitly
	fmt.Printf("%s\n", c)	// "100°C"
	fmt.Println(c)					// "100°C"
	fmt.Printf("%g\n", c)	// "100"; does not call String
	fmt.Println(float64(c))			// "100"; does not call String
}

func chapter2_6() {
	fmt.Println("\nchapter2_6 function call")
	//var ddd Celsius
	//var dd tempconv.Celsius
	//dd = 1.2

	var ddd int
	ddd = 10
	fmt.Printf("%d\n",ddd)

	var ddd2 tempconv.Celsius
	//var ddd2 tmppkg.Celsius
	ddd2 = 1.2
	fmt.Printf("%f\n",ddd2)
	//tmppkg.Celsius()


	var ddd3 tmppkg.Celsius
	ddd3 = 1.2
	fmt.Printf("%f\n",ddd3)
}

// init order
var a = b + c		// 3
var b = f3()		// 2
var c = 1			// 1
func f3() int { return  c + 1 }

func chapter2_6_2() {
	fmt.Println("\nchapter2_6_2 function call")
	fmt.Printf("a : %d\n",a)
	fmt.Printf("b : %d\n",b)
	fmt.Printf("c : %d\n",c)
}
