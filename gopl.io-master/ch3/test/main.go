
package main

import "fmt"

func main() {
	chapter3_1_1()
	chapter3_1_2()
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