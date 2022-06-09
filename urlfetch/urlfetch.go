/*******************************************************************************
 * urlfetch.go
 * combining chapters 1.2, 1.5, and 1.6 - cli args, fetching urls, and
 * concurrency (respectively).
 * authored by brianc2788@gmail.com
 * http://github.com/user5260
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

        if err !- nil {
            fmt.Fprintf(os.Stderr, "fetch read: %s (%v)\n", url, err)
            os.Exit(1)
        }
        fmt.Printf("%s", b)
    }
}