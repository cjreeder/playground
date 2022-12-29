package main

import (
	"fmt"

	Log "github.com/cjreeder/logging-library"
)

func main() {
	fmt.Println("Initial Print Statement using fmt")
	//Log.L.Infof("TESTING!")
	// Initialize the logger with the level
	LL := Log.NewLogger("info") // We create the logger from the logging-library
	LL.Log.Infof("More Output") // Logging out looks like this

}
