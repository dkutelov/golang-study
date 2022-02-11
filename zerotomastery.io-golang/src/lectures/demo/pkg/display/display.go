package display

// Can not use main func

import "fmt"

// Capitalised functions are exported from the package
func Display(msg string) {
	fmt.Println(msg)
}