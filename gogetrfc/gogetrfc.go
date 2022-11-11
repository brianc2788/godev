/*
******************************************************************************
gogetrfc.go
--------------------------------------------------------------------------------
IETF RFC (Request For Comment) Downloader.
--------------------------------------------------------------------------------
Powered by http://www.rfc-editor.org/
http://brianc2788.github.io/
******************************************************************************
*/
package main

import (
	"fmt"
	"os"
	//"net/http"
)

// target url format: http://www.rfc-editor.org/rfc/rfc[XXXX].txt (also available in pdf, etc.)
// rfc[XXXX] = user's desired rfc doc.

//func parseArgs()

// TODO: warmed-up by looping through & printing cli args.
func main() {
	if argArray := os.Args; len(argArray) > 1 {
		for i := 0; i < len(argArray); i++ {
			fmt.Printf("%s\n", argArray[i])
		}
	} else {
		fmt.Println("Didn't get any arguments...")
	}
}
