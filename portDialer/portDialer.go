/*
********************************************************************************
portDialer.go
-------------
Port scanner written in Go.
Currently, establishes a connection (completes TCP handshake).
Currently only checks port 80 (http) by default.
http://brianc2788.github.io
********************************************************************************
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

	//strings - better for readability?
	var (
		DestName       = CliArgs[1]
		PortNum        = "80"
		PortSep        = ":"
		PortSuffix     = PortSep + PortNum
		TransportLayer = "tcp"
	)

	// Dial it in.
	c, err := net.Dial(TransportLayer, (DestName + PortSuffix))
	if err != nil {
		fmt.Printf("Port %s is closed/filtered or host couldn't be contacted.", PortNum)
		panic(err)
	} else {
		c.Close()
		AddrAsString := c.RemoteAddr().String()
		AddrAsString = AddrAsString[:len(AddrAsString)-3]
		fmt.Printf("Established connection with %s (%s) on port %s.\n", DestName, AddrAsString, PortNum)
	}
}
