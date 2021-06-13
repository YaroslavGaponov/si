package main

import (
	"fmt"
	"github.com/YaroslavGaponov/si/lib"
)

func main() {

	s := si.New(2)
	s.Add("hello")
	s.Add("hi world")
	s.Add("hello world")
	s.Add("hallelujah world")

	for _, result := range s.Search([]string{"rl", "ll"}) {
		fmt.Println(result)
	}

}
