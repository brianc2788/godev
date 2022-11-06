/***********************************************************
 * quote-demo.go
 * -------------
 * Testing the rsc/quote. Demonstrate each of its methods.
 * http://brianc2788.github.io
***********************************************************/
package main
import (
    "fmt"
    "rsc.io/quote/v4"
)

func main() {
    qHello := quote.Hello()
    qGlass := quote.Glass()
    qGo := quote.Go()
    qOpt := quote.Opt()

    fmt.Printf("%s\n%s\n%s\n%s\n",qHello, qGlass, qGo, qOpt)
}