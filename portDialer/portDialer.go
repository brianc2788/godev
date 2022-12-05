/*
**********************************************************
  - portscan1.go

create a tcp socket and connect to "http://scanme.nmap.org",
aka "scanme.nmap.org:80".
http://brianc2788.github.io
**********************************************************
*/
package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	CliArgs := os.Args
	if len(CliArgs) <= 1 || len(CliArgs) > 2 {
		fmt.Printf("Need positional args: domain names.\n")
		//debug
		print("just one for now, please...")
		os.Exit(0)
	}

	DestIP := CliArgs[1]
	PortSuffix := ":80"
	TransportLayer := "tcp"

	// Dial it in.
	c, err := net.Dial(TransportLayer, (DestIP + PortSuffix))
	if err != nil {
		fmt.Printf("Error :%s\n", err)
	} else {
		fmt.Printf("Established connection with %s (%s).\n", DestIP, c.RemoteAddr())
	}
	c.Close()
}
