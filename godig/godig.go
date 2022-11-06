/***********************************************************
 * godig
 * -----
 * A version of the dig tool from gnu. Sort of.
 *
 http://brianc2788.github.io
***********************************************************/
package main
import (
	"fmt"
	"os"
	"github.com/miekg/dns"
)

func sendQuery(domainName string) {
	var msg dns.Msg
	fqdn := dns.Fqdn(domainName)
	msg.SetQuestion(fqdn,dns.TypeA)
	dns.Exchange(&msg, "8.8.8.8:53")
	fmt.Println(msg)
}

func main() {
	// Loop through args and query each.
	// Print output (still need to decompress, parse, etc.)
	// Answers can be seen via tcpdump/packet capture.
	for i := 1; i < len(os.Args); i++ {
		sendQuery(os.Args[i])
	}
}
