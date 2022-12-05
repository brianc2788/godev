/********************************************
* hellogo.go
* testing several print functions.
* can you tell that i struggle to come up
* with arbitrary text?
********************************************/
package main

import "fmt"

func main() {
	// locals
	worldstr := "sekai"
	//I can't get the book's kanji to work, so here's the romanji. yeah, I know a little.
	goauthor1 := "Ken Thompson"

	//alt declaration
	var (
		var1 string
		var2 string
	)
	// assignment
	var1 = "thing1"
	var2 = "thing2"

	fmt.Println("Hello,", goauthor1, "!") //has a separator
	fmt.Printf("I mean, jello, %s!\n", worldstr)
	fmt.Println("I mean...\nHello, Println, Printf, Print, et al.")
	fmt.Print("Print() just prints the stuff you give it; no separators,",
		"formatting, etc.\nsee? no newline.\n",
		"Oh, actually, it does do the bslash esc newlines... huh.\n")
	fmt.Printf("more debug:\n%s - %s\n", var1, var2)

	print("Just ignore this binary and whatever directory you found it in...\n")

	fmt.Print("Exiting.......")
}
