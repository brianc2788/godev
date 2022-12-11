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
	"log"
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
		go TcpConnect(dname, args[n], &wgroups)
	}

	wgroups.Wait()
}

func spinner() {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("*Scanning [%c]\r", r) // \r won't work on windows?
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func TcpConnect(a string, p string, j *sync.WaitGroup) {
	to, terr := time.ParseDuration("10s")
	defer j.Done()
	if terr != nil {
		log.Fatalf("%s\n", terr)
	}
	c, err := net.DialTimeout("tcp", a+":"+p, to)
	if err != nil {
		fmt.Printf("\r%s - port %s filtered|closed\n", a, p)
	} else {
		fmt.Printf("\r%s - port %s open\n", a, p)
		/* Undesired effect. Move somewhere else?
		ip4 := c.RemoteAddr().String()
		ip4 = ip4[:((len(ip4)-len(p))-1)]
		fmt.Printf("Finished scanning %s (%s)\n", a, ip4)
		*/
		c.Close()
	}
}
