package main

import "fmt"

func main() {
	// declare some local variables.
	worldstr := "world"
	UltimateHuman := "Ken Thompson"
	var1,var2 := "thing1","thing2"
	
	//bunches of ways to print stuff...
	//every beginner tutorial wants to use Println for some reason.
	//however, there's apparently a print function for any situation...
	fmt.Println("Hello,", UltimateHuman, "!") //note: fmt inserts a default seperator.
	fmt.Printf("I mean, jello, %s!\n", worldstr)
	fmt.Println("I mean...\nHello, Println, Printf, Print, et al.")
	fmt.Print("Print() just prints the stuff you give it; no separators, formatting, etc.\nsee? no newline.")
	fmt.Printf("\nmore debug:\n%s - %s\n",var1,var2)
	fmt.Print("Exiting.......")
}
