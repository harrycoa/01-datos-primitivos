package mypackage

import "fmt"

// CarPublic struct car con acceso publico
type CarPublic struct {
	Brand string
	Year  int
	Color string
	Model string
}

type carPrivate struct {
	brand string
	year  int
	color string
	model string
}

// PrintMessage
func PrintMessage() {
	fmt.Println("Hello from package mypackage")
}

// printMessage
func printMessage() {
	fmt.Println("Hello from package private mypackage")
}
