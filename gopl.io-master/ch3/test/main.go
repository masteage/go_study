
package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"strconv"
	"time"
	"unicode/utf8"
)

//⁰¹²³⁴⁵⁶⁷⁸⁹⁺⁻⁼⁽⁾ⁿ
func main() {
	chapter3_1_1()
	chapter3_1_2()
	chapter3_1_3()
	chapter3_2()
	chapter3_3()
	chapter3_5()
	chapter3_5_3_01()
	chapter3_5_3_02()
	chapter3_5_5()
	chapter3_6()
	chapter3_6_1()
	chapter3_6_2()
}

func chapter3_1_1() {
	fmt.Println("\nchapter3_1_1 function call\n")

	// %
	fmt.Println("\n'%' test\n")
	fmt.Printf("-5%%3 : %v\n",-5%3)
	fmt.Printf("5%%-3 : %v\n",5%-3)
	fmt.Printf("-5%%-3 : %v\n",-5%-3)

	// /
	fmt.Println("\n'/' test \n")
	fmt.Printf("5/4 : %v\n",5/4)
	fmt.Printf("5.0/4 : %v\n",5.0/4)
	fmt.Printf("5/4.0 : %v\n",5/4.0)

	// overflow
	fmt.Println("\noverflow test ")
	var u uint8 =  255			// 0 ~ 255 (0 ~ 2ⁿ-1)
	fmt.Println(u, u+1, u*u)	// 255*255 = 65025 = 256*254 + 1
								// (2ⁿ-1) * (2ⁿ-1) = (2ⁿ*2ⁿ) - (2*2ⁿ) + 1 = 2ⁿ(2ⁿ - 2) + 1 = 1
	var i int8 = 127			// -128 ~ 127 (-2ⁿ⁻¹ ~ 2ⁿ⁻¹-1)
	fmt.Println(i, i+1, i*i)
}

func chapter3_1_2() {
	fmt.Println("\nchapter3_1_2 function call\n")
	var x uint8 = 1<<1 | 1<<5
	var y uint8 = 1<<1 | 1<<2

	fmt.Printf("x\t\t%08b\n",x)
	fmt.Printf("y\t\t%08b\n",y)

	fmt.Printf("x&y\t\t%08b\n",x&y)
	fmt.Printf("x|y\t\t%08b\n",x|y)

	fmt.Printf("x^y\t\t%08b\n",x^y)
	fmt.Printf("x&^y\t%08b\n",x&^y)

	fmt.Printf("x<<1\t%08b\n",x<<1)
	fmt.Printf("x>>1\t%08b\n",x>>1)

	// x
	for i := uint(0); i < 8 ; i++ {
		if x&(1<<i) != 0 {
			fmt.Println(i)
		}
	}

	// y
	for i := uint(0); i < 8 ; i++ {
		if y&(1<<i) != 0 {
			fmt.Println(i)
		}
	}
}

func chapter3_1_3() {
	fmt.Println("\nchapter3_1_3 function call\n")
	var apples int32 = 1
	var oranges int16 = 2
	//var compote int = apples + oranges	// compile error
	var compote int = int(apples) + int(oranges)
	fmt.Printf("compote is %d\n",compote)

	f := 3.14
	i := int(f)
	fmt.Println(f,i)
	f = 1.99
	fmt.Println(int(f))

	f2 := 1e100
	i2 := int(f2)
	fmt.Println(f2,i2)

	o := 0666
	fmt.Printf("%d %[1]o %#[1]o\n",o)
	x := int64(0xdeadbeef)
	fmt.Printf("%d %[1]x %#[1]X\n",x)

	ascii := 'a'
	unicode := '©'
	newline := '\n'

	fmt.Printf("%d %[1]c %[1]q\n",ascii)
	fmt.Printf("%d %[1]c %[1]q\n",unicode)
	fmt.Printf("%d %[1]q\n",newline)
}

func chapter3_2() {
	fmt.Println("\nchapter3_2 function call\n")

	fmt.Println(math.MaxInt8)
	fmt.Println(math.MaxInt16)
	fmt.Println(math.MaxInt32)
	fmt.Println(math.MaxInt64)
	fmt.Println(math.MaxFloat32)
	fmt.Println(math.MaxFloat64)

	var f float32 = 1 << 24
	fmt.Println(f)
	fmt.Println(f == f+1)

	for x:= 0; x < 8; x++ {
		fmt.Printf("x = %d e(x) = %8.3f\n",x,math.Exp(float64(x)))
	}

	var z float64
	fmt.Println(z, -z , 1/z , -1/z , z/z)

	nan := math.NaN()
	fmt.Println(nan == nan, nan < nan, nan > nan)
}

func chapter3_3() {
	fmt.Println("\nchapter3_3 function call\n")

	var x complex128 = complex(1,2)
	var y complex128 = complex(3,4)
	fmt.Println(x*y)
	fmt.Println(real(x*y))
	fmt.Println(imag(x*y))

	z := 1 + 2i
	fmt.Println(z)

	fmt.Println(cmplx.Sqrt(-1))
}

func chapter3_5() {
	fmt.Println("\nchapter3_5 function call\n")
	s := "hello, world"
	fmt.Println(s[0],s[5])

	// panic
	//c := s[len(s)]
	//fmt.Println(c)
	fmt.Println(s[0:5])
	fmt.Println(s[:5])
	fmt.Println(s[7:])
	fmt.Println(s[:])
	fmt.Println("goodbye" + s[5:])

	s2 := "ddd"
	//s2[0] = "d"
	fmt.Println(s2)
	//t := "\123"
	//fmt.Println(t)
}

func chapter3_5_3_01() {
	fmt.Println("\nchapter3_5_3_01 function call\n")

	fmt.Println("世界")
	fmt.Println("\xe4\xb8\x96\xe7\x95\x8c")
	fmt.Println("\u4e16\u754c")
	fmt.Println("\U00004e16\U0000754c")

	s := "hello, 世界"
	fmt.Println(len(s))
	fmt.Println(utf8.RuneCountInString(s))
	for i := 0; i< len(s); {
		r, size := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%d\t%d\t%c\n",i,size,r)
		i += size
	}

	for i, r := range s {
		fmt.Printf("%d\t%q\t%d\n",i,r,r)
	}
	n := 0
	for range s {
		n++
	}
	fmt.Printf("n : %d\n",n)
	fmt.Printf("count : %d\n",utf8.RuneCountInString(s))
}

func chapter3_5_3_02() {
	fmt.Println("\nchapter3_5_3_02 function call\n")

	s := "한국어"
	fmt.Printf("% X\n",s)
	r := []rune(s)
	fmt.Printf("%x\n",r)

	fmt.Println(string(s))
	fmt.Println(string(r))

	fmt.Println(string(0x4eac))
	fmt.Println(string(1234567))
	//fmt.Println(rune(string(1234567)))
}

func chapter3_5_5() {
	fmt.Println("\nchapter3_5_5 function call\n")

	x := 123
	y := fmt.Sprintf("%d",x)
	fmt.Println(y,strconv.Itoa(x))

	fmt.Println(strconv.FormatInt(int64(x),2))
	fmt.Println(strconv.FormatInt(int64(x),10))

	fmt.Println(fmt.Sprintf("x=%b",x))

	x2,err := strconv.Atoi("123")
	y2,err := strconv.ParseInt("123",10,64)
	if err != nil {
		return
	}
	fmt.Println(x2,y2)
}

func chapter3_6() {
	fmt.Println("\nchapter3_6 function call\n")

	const myLen = 3
	var myVar [myLen]byte
	fmt.Println(len(myVar))

	const noDelay time.Duration = 0
	const timeout = 5 * time.Minute
	fmt.Printf("%T %[1]v\n",noDelay)
	fmt.Printf("%T %[1]v\n",timeout)
	fmt.Printf("%T %[1]v\n",time.Minute)

	const (
		a = 1
		b
		c = 2
		d
	)
	fmt.Println(a,b,c,d)
}

func chapter3_6_1() {
	fmt.Println("\nchapter3_6_1 function call\n")

	type MYENUM = int
	const (
		A MYENUM = iota
		B
		C
	)
	fmt.Println(A,B,C)

	type MY_FLAG = uint
	const (
		F_A MY_FLAG = 1 << iota
		F_B
		F_C
		F_D
	)
	fmt.Println(F_A,F_B,F_C,F_D)
}

func chapter3_6_2() {
	fmt.Println("\nchapter3_6_2 function call\n")

	var x float32 = math.Pi
	var y float64 = math.Pi
	var z complex128 = math.Pi
	fmt.Println(x,y,z)

	const Pi64 float64 = math.Pi
	var x1 float32 = float32(Pi64)
	var y1 float64 = Pi64
	var z1 complex128 = complex128(Pi64)
	fmt.Println(x1,y1,z1)

	var f float64 = 212
	fmt.Println((f-32) * 5 / 9)
	fmt.Println(5 / 9 * (f-32))
	fmt.Println(5.0 / 9.0 * (f-32))

	fmt.Println(5 / 9)
	fmt.Println(5.0 / 9.0)

}