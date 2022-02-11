package main

import (
	"coursecontent/demo/pkg/display"
	"coursecontent/demo/pkg/msg"
)

// Package should be in a separate folder
// Main launches the program

func main() {
	msg.Hi()
	display.Display("Hello from display")
	msg.Exciting("an exciting message")
}