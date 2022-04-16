package main

import (
    "fmt"
    "time"
    "regexp"
)

func main() {
    timere := regexp.MustCompile(`\d\d:\d\d:\d\d`)
    currentTime := time.Now()
    fmt.Println("string:",currentTime)
    fmt.Println("filtered:",timere.FindString(string(`currentTime`)))
}

