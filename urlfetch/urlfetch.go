/*******************************************************************************
 * urlfetch.go
 * -----------
 * combining chapters 1.2, 1.5 - cli args, fetching urls.
 * Right now, just takes a URL (with schema) and returns the html.
 * http://brianc2788.github.io
*******************************************************************************/
package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "os"
)

func main() {
	for _,url := range os.Args[1:] {
        resp,err := http.Get(url)
        
        if err != nil {
            fmt.Fprintf(os.Stderr, "fetch error: %v\n", err)
            os.Exit(1)
        }

        b,err := ioutil.ReadAll(resp.Body)
        resp.Body.Close()

        if err != nil {
            fmt.Fprintf(os.Stderr, "fetch read: %s (%v)\n", url, err)
            os.Exit(1)
        }
        fmt.Printf("%s", b)
    }
}
