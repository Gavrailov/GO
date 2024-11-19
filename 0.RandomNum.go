//fmt is used for formating and print

// In Go, a name is exported if it begins with a capital letter. For example, Pizza is an exported name, as is Pi, which is exported from the math package.

//pizza and pi do not start with a capital letter, so they are not exported.

package main

import (
	"fmt"
	"math/rand"
)

func main0() {
	fmt.Println("My favorite number is", rand.Intn(10))
}
