package main

import (
	"fmt"
	"unsafe"
)

func main() {
	s := "eegoo"

	fmt.Printf("%c\n", s[1])

	p := unsafe.Pointer(&s)

	t1 := (*[]byte)(p)

	fmt.Println(string(*t1))
	t := *t1
	t[1] = 'c'

	fmt.Println(string(t))

}
