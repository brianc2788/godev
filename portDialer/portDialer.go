/*
portDialer version 2.1.0
------------------------
Added concurrent goroutines.
TODO: Parse args for ports and ranges.

	Improve concurrent subroutines.

http://brianc2788.github.io/
*/
package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	args := os.Args
	if len(args) < 3 {
		fmt.Printf("Need domain & port(s)\n")
		os.Exit(0)
	}
	dname := args[1]

	go spinner()

	// Loop through args[2:] (ports)
	for n := 2; n < len(args); n++ {
		go checkPort(dname, args[n])
	}

	// wait for goroutines
	time.Sleep(11 * time.Second)
}

func spinner() {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("%c\r", r)
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func checkPort(a string, p string) {
	to, terr := time.ParseDuration("10s")
	if terr != nil {
		panic(terr)
	}
	c, err := net.DialTimeout("tcp", a+":"+p, to)
	if err != nil {
		fmt.Printf("%s - port %s filtered|closed\n", a, p)
	} else {
		fmt.Printf("%s - port %s open\n", a, p)
		/* Printed out ipv4, but didn't come out right. Move somewhere else?
		ip4 := c.RemoteAddr().String()
		ip4 = ip4[:((len(ip4)-len(p))-1)]
		fmt.Printf("Finished scanning %s (%s)\n", a, ip4)
		*/
		c.Close()
	}
}
