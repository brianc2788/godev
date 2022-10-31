/***********************************************************
 * godig
 *
 * A version of the dig tool from gnu. Sort of.
 *
 http://brianc2788.github.io
***********************************************************/
package main
import (
	"fmt"
	"github.com/miekg/dns"
)

func main() {
	var msg dns.MSG
	fqdn := Fqdn("stacktitan.com")
	msg.SetQuestion(fqdn,dns.TypeA)
	dns.Exchange(&msg, "8.8.8.8:53")
}
