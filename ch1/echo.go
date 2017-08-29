package main

import (
	"fmt"
	"os"
	"strings"
)

func linebreak() {
	fmt.Println("-------------------------------------------------------------")
}

func echo1() {
	fmt.Println("Echo1:")
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}

func echo2() {
	fmt.Println("Echo2:")
	var s, sep string
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

func echo3() {
	fmt.Println("Echo3:")
	fmt.Println(strings.Join(os.Args[1:], " "))
}

func main() {
	echo1()
	linebreak()
	echo2()
	linebreak()
	echo3()
}
