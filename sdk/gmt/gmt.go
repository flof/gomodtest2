package gmt

import "fmt"

func Create(debug bool) string {
	return fmt.Sprintf("Version 2.0.2: debug %v", debug)
}
