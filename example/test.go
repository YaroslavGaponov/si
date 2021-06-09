package main

import (
	"github.com/YaroslavGaponov/si"
	"fmt"
)

func main() {

	s := si.NewSearchIndex(3)
	s.add("hello")
	s.add("hi world")
	s.add("hello world")

	var prompt = "rld"
	result := s.search(prompt)

	fmt.Printf("search <%s>\n", prompt)
	for i := 0; i < len(result); i++ {
		fmt.Printf("<%s>\n", result[i])
	}
	fmt.Printf("total %d\n", len(result))

}
