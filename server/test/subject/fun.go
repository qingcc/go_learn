package main

import (
	"fmt"
)

func change(s ...string) {
	//fmt.Println("in function: s:", &s[0])
	s[0] = "Go"
	s = append(s, "playground")
	fmt.Println(s)
	//fmt.Println("in function: s:", &s[0])
}

func main() {
	welcome := []string{"hello", "world"}
	//fmt.Println("out of function: s:", &welcome[0])
	change(welcome...)
	//fmt.Println("after change: s:", &welcome[0])
	fmt.Println(welcome)
}
