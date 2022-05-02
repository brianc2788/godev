/********************************************
* hellogo.go
* Slightly more printing than your average
* "HelloWorld" code.
********************************************/
package main

import (
	"fmt"
	//"rsc.io/quote" // cool quote i found in the Go Tour thing.
	"https://github.com/rsc/quote/blob/v4.0.1/quote.go"
)

func main() {
	// declare some local variables.
	worldstr := "world"
	UltimateHuman := "Ken Thompson"

	// alternate declaration;
	// as opposed to "name := value".
	var (
		var1 string
		var2 string
	)
	// assignment
	var1 = "thing1"
	var2 = "thing2"
	
	fmt.Println("Hello,", UltimateHuman, "!") //has a separator
	fmt.Printf("I mean, jello, %s!\n", worldstr)
	fmt.Println("I mean...\nHello, Println, Printf, Print, et al.")
	fmt.Print("Print() just prints the stuff you give it; no separators, formatting, etc.\nsee? no newline.\nOh, actually, it does do the bslash esc newlines... huh.\n")
	fmt.Printf("more debug:\n%s - %s\n",var1,var2)

	// another alternate declaration...
	var var3 = quote.Go()
	fmt.Print(var3)

	fmt.Print("Exiting.......")
}
