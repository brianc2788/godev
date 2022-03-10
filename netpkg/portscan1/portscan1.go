/***********************************************************
create a tcp socket and connect to "http://scanme.nmap.org",
aka "scanme.nmap.org:80".
brianc2788@gmail.com
github.com/user5260/go/netpkg/portscan1
------------------------------------------------------------
TODO:
	print errors
	get url & port from user.
	create a jupyter notebook for creating socket
		connections with go.
***********************************************************/
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

	// Dial(network, address)
	// 
	_, err := net.Dial(socket,host)
	if err == nil {
		fmt.Println("Connected to",host)	// sep defaults to ' '.
	}
}