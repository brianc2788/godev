/*
**********************************************************
* godig.go
* --------------------------------------------------------
* Clone of dig (not really) from gnu written in Go.
* TODO: Reduce bloat.
* - Parse, print answer(s).
* - Add TLD challenge to processName().
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
		nameIn := argList[userArg]
		fqname := processName(nameIn)
		rMsg := new(dns.Msg)

		rMsg, merr := getRecordA(fqname)
		if merr != nil {
			panic(merr)
		}

		fmt.Printf("### INPUT ###\n")
		fmt.Printf("name: %s\n", argList[userArg])
		fmt.Printf("fqdn: %s\n", fqname)
		for _, q := range rMsg.Question {
			fmt.Printf("payload: %s\n", q.Name)
		}
		fmt.Printf("\n")
	}
}

/* Check user's domain names. */
func processName(TestName string) string {
	if !dns.IsFqdn(TestName) {
		return dns.Fqdn(TestName)
	}
	return TestName
}

/* Returns the IPv4 of DNS server; with separator & port. */
func getDnsAddr() string {
	fullAddr := DNS_GOOGLE_ADDR + DNS_SEP + DNS_PORT
	return fullAddr
}

/* Resolve names via DNS, print to stdout. */
func getRecordA(dn string) (*dns.Msg, error) {
	var nMsg dns.Msg
	nMsg.SetQuestion(dn, dns.TypeA)
	_, err := dns.Exchange(&nMsg, getDnsAddr())

	if err != nil {
		return nil, fmt.Errorf("Error during dns exchange.\nRuntime Error:\n%s", err)
	}

	return &nMsg, nil
}
