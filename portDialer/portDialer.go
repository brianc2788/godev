/*
********************************************************************************
portDialer version 2.1.0
------------------------
Concurrent port scanner. Open port indicates completion of a TCP connection(!).
TODO: -Print help/usage.

	      -Parse args for ports; comma-seperated & hyphenated-range formats.
		  -Emulate a nmap stealth scan by not completing TCP 3-way handshake (?).

http://brianc2788.github.io/
********************************************************************************
*/
package main

import (
	"fmt"
	"net"
	"os"
	"sync"
	"time"
)

func main() {
	args := os.Args
	if len(args) < 3 {
		fmt.Printf("Enter domain name & white-space seperated port(s).\n")
		os.Exit(0)
	}
	dname := args[1]
	var wgroups sync.WaitGroup

	go spinner()

	// Loop through args[2:] (ports)
	for n := 2; n < len(args); n++ {
		wgroups.Add(1)
		go checkPort(dname, args[n], &wgroups)
	}

	wgroups.Wait()
}

func spinner() {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("*Scanning [%c]\r", r)
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func checkPort(a string, p string, j *sync.WaitGroup) {
	to, terr := time.ParseDuration("10s")
	defer j.Done()
	if terr != nil {
		panic(terr)
	}
	c, err := net.DialTimeout("tcp", a+":"+p, to)
	if err != nil {
		fmt.Printf("\r%s - port %s filtered|closed\n", a, p)
	} else {
		fmt.Printf("\r%s - port %s open\n", a, p)
		/* Printed out ipv4, but didn't come out right. Move somewhere else?
		ip4 := c.RemoteAddr().String()
		ip4 = ip4[:((len(ip4)-len(p))-1)]
		fmt.Printf("Finished scanning %s (%s)\n", a, ip4)
		*/
		c.Close()
	}
}
