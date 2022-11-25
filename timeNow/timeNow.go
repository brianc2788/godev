/**********************************
* timeNow.go
* http://brianc2788.github.io/
**********************************/
package main

import (
    "fmt"
    "time"
    "regexp"
)

func main() {
    /* Time Regex
       Doesn't match "0" prefix for the hours
       but does so for minutes & seconds.
       Truncates seconds to 2 digits.
    */
    timere := regexp.MustCompile(`1?[1-9]?:\d{1,2}:\d{1,2}`)
    currentTime := time.Now()
    filteredTime := timere.FindString(currentTime.String())
    fmt.Println("string:",currentTime)
    fmt.Println("filtered:", filteredTime)
}

