package main

import "fmt"

type sliceVal []string

func (s sliceVal) String() string {
	var str string
	for _, i := range s {
		str += fmt.Sprintf("%s\n", i)
	}
	return str
}
