package main

import (
	"github.com/YaroslavGaponov/si/lib"
	"fmt"
)

func main() {

	s := si.New(3)
	s.Add("hello")
	s.Add("hi world")
	s.Add("hello world")

	var prompt = "orld"
	result := s.Search(prompt)

	fmt.Printf("search <%s>\n", prompt)
	for i := 0; i < len(result); i++ {
		fmt.Printf("<%s>\n", result[i])
	}
	fmt.Printf("total %d\n", len(result))

}
