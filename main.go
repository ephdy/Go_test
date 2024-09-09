package main

import (
	pp "awesomeProject/package"
	"fmt"
)

func main() {
	s := "gopher"
	fmt.Println("Hello and welcome, %s!", s)

	for i := 1; i <= 5; i++ {
		fmt.Println("i =", 100/i)
	}

	pp.F2()
}

func init() {
	fmt.Println("调用了main-init")
}
