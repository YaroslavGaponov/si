Search Index
===========


# Demo

Golang playground [here](https://play.golang.org/p/WxYsnoe-lXm)

## Test

```golang
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

	fmt.Println(
		s.Search([]string{"orl", "ll"}),
	)

}
```

## Result

```sh
[hello world]
```