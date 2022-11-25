/*
**********************************************************
  - godig
  - -----
  - A version of the dig tool from gnu. Sort of.
  - Inspired by "Black Hat Go" by Tom Steele, Chris Patten, and Dan Kottmann.
    *
  - TODO:-Parse and, if necessary, decompress dns response.
  - -Impl. arg parser with switches for options.
  - -Opts include dns server, query type, etc.
  - -Impl. other dns query types, services.
  - -DNSec?
  - -Continue improving code readability & optimize.
    http://brianc2788.github.io

**********************************************************
*/
package main

import (
	"fmt"
	"os"

	"github.com/miekg/dns"
)

/* Constants - Port & DNS Server Addresses (ipv4). */
const DNS_SEP = ":"
const DNS_PORT = "53"
const DNS_GOOGLE_ADDR = "8.8.8.8"
const DNS_CLOUDFLARE_ADDR = "1.1.1.1"
const DNS_LOCAL_ADDR = "192.168.1.1"

func getDnsAddr() string {
	fullAddr := DNS_GOOGLE_ADDR + DNS_SEP + DNS_PORT
	return fullAddr
}

/* Resolve names via DNS, print to stdout. */
func sendQuery(domainName string) dns.Msg {
	var msg dns.Msg
	fqdn := dns.Fqdn(domainName) // Don't need trailing period.
	msg.SetQuestion(fqdn, dns.TypeA)
	dns.Exchange(&msg, getDnsAddr())

	return msg
}

/* PTR Record queries. Not yet implemented.
------------------------------------------------------------
   dig example:
   (Querying ipv4 address 94.177.245.177)
   user@localhost$ dig 177.245.177.94.in-addr.arpa PTR
   Note: ipv6 format: [addr].ipv6.arpa
------------------------------------------------------------

func getRecordPTR(address string) dns.Msg {
	//TODO: Reverse DNS lookup.
}
*/

func main() {
	argList := os.Args
	argCount := len(argList)

	/* Loop through args with sendQuery(). */
	for userArg := 1; userArg < argCount; userArg++ {
		response := sendQuery(argList[userArg])
		fmt.Printf("%s is at %s\n", argList[userArg], response.Answer)
	}
}
