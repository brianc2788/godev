/***********************************************************
 * portscan1.go
create a tcp socket and connect to "http://scanme.nmap.org",
aka "scanme.nmap.org:80".
brianc2788@gmail.com
github.com/user5260/go
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
	// Returns TWO variables; being assigned to _ and err (error string).
	_, err := net.Dial(socket,host)
	if err != nil {
		fmt.Printf("Error :%s\n",err)
	} else {
		fmt.Printf("Successful connection to %s.\n",host)
	}
}
