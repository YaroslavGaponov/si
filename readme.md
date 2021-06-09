Search Index
===========


# Demo

## Test
```
package main

import (
	"fmt"
)

func main() {

	si := NewSearchIndex(3)
	si.add("hello")
	si.add("hi world")
	si.add("hello world")

	var prompt = "rld"
	result := si.search(prompt)

	fmt.Printf("search <%s>\n", prompt)
	for i := 0; i < len(result); i++ {
		fmt.Printf("<%s>\n", result[i])
	}
	fmt.Printf("total %d\n", len(result))

}
```

## Result
```terminal
search <rld>
<hi world>
<hello world>
total 2
```