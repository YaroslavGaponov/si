Search Index
===========


# Demo

## Test

```golang
package main

import (
	"github.com/YaroslavGaponov/si"
	"fmt"
)

func main() {

	s := si.NewSearchIndex(3)
	s.Add("hello")
	s.Add("hi world")
	s.Add("hello world")

	var prompt = "rld"
	result := s.Search(prompt)

	fmt.Printf("search <%s>\n", prompt)
	for i := 0; i < len(result); i++ {
		fmt.Printf("<%s>\n", result[i])
	}
	fmt.Printf("total %d\n", len(result))

}
```

## Result

```sh
search <rld>
<hi world>
<hello world>
total 2
```