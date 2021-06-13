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

	var prompt = []string{"rl", "ll"}
	result := s.Search(prompt)

	fmt.Printf("search <%s>\n", prompt)
	for i := 0; i < len(result); i++ {
		fmt.Printf("<%s>\n", result[i])
	}
	fmt.Printf("total %d\n", len(result))

}
