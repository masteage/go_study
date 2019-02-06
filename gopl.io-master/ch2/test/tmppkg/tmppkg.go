// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 41.

//!+

package tmppkg

import "fmt"

type Celsius float64
type Fahrenheit float64

// not able : myType -> MyType
//type myType int
type MyType int

// not able : myFunc -> MyFunc
//func (t MyType) myFunc() string { return fmt.Sprintf("Value(%d)", t) }
func (t MyType) MyFunc() string { return fmt.Sprintf("Value(%d) - MyFunc", t) }
func (t MyType) String() string { return fmt.Sprintf("Value(%d) - String", t) }
//!-
