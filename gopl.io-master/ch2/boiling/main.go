// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 29.
//!+

// Boiling prints the boiling point of water.


// https://golang.org/pkg/fmt/#hdr-Printing

// https://golang.org/ref/spec#Keywords
// https://golang.org/ref/spec#Constants
// https://golang.org/ref/spec#Types
// https://golang.org/ref/spec#Built-in_functions


// package declaration
package main

// import package
import "fmt"

// package variable
const boilingF = 212.0

func main() {

	// local variable
	var f = boilingF
	var c = (f - 32) * 5 / 9

	fmt.Printf("boiling point = %g°F or %g°C\n", f, c)
}

//!-
