package main

import (
	"fmt"
	"sizedlist"
)

var long string = "A string woo"

type LenString string

func (l LenString) Len() int {
	return len(l)
}

func m(x int) LenString {
	return LenString(long)
}

func main() {
	//test count list
	c := sizedlist.NewCountList(3)
	c.Append(m(1), m(2), m(3))
	for i := 0; i < 1e3; i++ {
		c.Append(m(i))
	}
	fmt.Println(c.Data)

	//test byte list
	b := sizedlist.NewBytesList(2000)
	for i := 0; i < 1e3; i++ {
		b.Append(m(i))
	}
	fmt.Println(b.Data)
	fmt.Println(b.Len())
}
