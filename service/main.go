package main

import (
	"fmt"

	"github.com/flof/gomodtest2/sdk/gmt/v2"
)

func main() {
	v := gmt.Create(true)
	fmt.Println(v)
}
