/*
use net module to create 
*/
package main
import (
	"fmt"
	"net"
)

func main() {
	// sock stream (tcp) or sock dgram (udp).
	// host is host domain & port.
	var socket = "tcp"
	var host = "scanme.nmap.org:80"

	_, err := net.Dial(socket,host)
	if err == nil {
		fmt.Println("Connected to",host)	// sep defaults to ' '.
	}
	/* TODO:
		print error
		close socket
	*/
}