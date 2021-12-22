package main

import (
	"fmt"

	"github.com/flof/gomodtest2/sdk/gmt/v3"
)

func main() {
	v := gmt.Create(true, false)
	fmt.Println(v)
}
