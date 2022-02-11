package msg

// Should import from the root of top mod file
import "coursecontent/demo/pkg/display"


func Hi() {
	//type the func and vscode will import
	display.Display("Hi")
}