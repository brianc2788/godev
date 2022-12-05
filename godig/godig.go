/*
********************************************************************************
* godig.go
* ------------------------------------------------------------------------------
* Clone of BIND's digutil (not really) written in Go.
* TODO: Reduce bloat.
* - Parse, print answer(s).
* - Use TTL to determine if domain was cached.
* - Check TLD in processName().
* - Implement reverse lookups (PTR records); CLI switch for rec types.
* - User Server input
* - Goroutines (concurrency)
http://brianc2788.github.io/
********************************************************************************
*/
package main

import (
	"fmt"
	"os"

	"github.com/miekg/dns"
)

/* Constants for domain server options. */
const (
	RESOLVER_URL_SEP         = ":"
	RESOLVER_PORT            = "53"
	RESOLVER_IPV4_GOOGLE     = "8.8.8.8"
	RESOLVER_IPV4_CLOUDFLARE = "1.1.1.1"
	RESOLVER_IPV4_LOCAL      = ""
)

/* MAIN */
func main() {
	argVars := os.Args
	argCount := len(argVars)

	/* For each domain arg, process name, exchange dgrams, output info. */
	for nArg := 1; nArg < argCount; nArg++ {
		nameIn := argVars[nArg]
		fqname := processName(nameIn)
		rMsg := new(dns.Msg)

		rMsg, merr := getRecordA(fqname)
		if merr != nil {
			panic(merr)
		}
		printQuestions(rMsg)
		//printAnswers(rMsg)
	}
}

/* Check user's domain names. */
func processName(dname string) string {
	if !dns.IsFqdn(dname) {
		return dns.Fqdn(dname)
	}
	return dname
}

/* Concatonate domain name variables. */
func catDnsAddr() string {
	fullAddr := RESOLVER_IPV4_GOOGLE + RESOLVER_URL_SEP + RESOLVER_PORT
	return fullAddr
}

func printQuestions(msgp *dns.Msg) {
	for n, q := range msgp.Question {
		fmt.Printf("payload %d: %s\n", n+1, q.Name)
	}
	fmt.Printf("\n")
}

func printAnswers(msgp *dns.Msg) {
	/*
		for n, a := range msgp.Answer {
			fmt.Printf("answer %d: %s\n", n, a)
		}
		fmt.Printf("\n")
	*/
}

/* Resolve names via DNS, print to stdout. */
func getRecordA(dn string) (*dns.Msg, error) {
	var nMsg dns.Msg
	nMsg.SetQuestion(dn, dns.TypeA)
	_, err := dns.Exchange(&nMsg, catDnsAddr())

	if err != nil {
		return nil, fmt.Errorf("Error during dns exchange.\nRuntime Error:\n%s", err)
	}

	return &nMsg, nil
}
