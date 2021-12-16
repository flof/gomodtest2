package gmt

import "fmt"

func Create(debug bool) string {
	return fmt.Sprintf("Version 1.0.0: debug %v", debug)
}
