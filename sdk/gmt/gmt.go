package gmt

import "fmt"

func Create(debug, verbose bool) string {
	return fmt.Sprintf("Version 3.0.0: debug %v, verbose: %v", debug, verbose)
}
