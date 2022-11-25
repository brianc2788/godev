/*
**********************************************************
* godig.go
* --------------------------------------------------------
* dig wannabe written in Go.
* TODO:
* - Parse, print answer(s).
* - PTR records for reverse lookup (ipv4 -> name).
http://brianc2788.github.io/
**********************************************************
*/
package main

import (
	"fmt"
	"os"

	"github.com/miekg/dns"
)

const DNS_SEP = ":"
const DNS_PORT = "53"
const DNS_GOOGLE_ADDR = "8.8.8.8"
const DNS_CLOUDFLARE_ADDR = "1.1.1.1"
const DNS_LOCAL_ADDR = "192.168.1.1" /* Leave empty for dns mod? */

/* MAIN */
func main() {
	argList := os.Args
	argCount := len(argList)

	/* Loop through args with sendQuery(). */
	for userArg := 1; userArg < argCount; userArg++ {
		name := argList[userArg]
		fqname := dns.Fqdn(name)
		rMsg := getRecordA(name)

		fmt.Printf("### INPUT ###\n")
		fmt.Printf("name: %s\n", argList[userArg])
		fmt.Printf("fqdn: %s\n", fqname)
		for _, q := range rMsg.Question {
			fmt.Printf("payload: %s\n", q.Name)
		}
		fmt.Printf("\n")
	}
}

/* Returns the IPv4 of DNS server; with separator & port. */
func getDnsAddr() string {
	fullAddr := DNS_GOOGLE_ADDR + DNS_SEP + DNS_PORT
	return fullAddr
}

/* Resolve names via DNS, print to stdout. */
func getRecordA(domainName string) *dns.Msg {
	var nMsg dns.Msg
	fqdn := dns.Fqdn(domainName) // Don't need trailing period.
	nMsg.SetQuestion(fqdn, dns.TypeA)
	_, err := dns.Exchange(&nMsg, getDnsAddr())

	if err != nil {
		fmt.Printf("Problem exchanging datagrams; Aborting.\nError/StackTrace:\n")
		panic(err)
	}

	return &nMsg
}
