/*
******************************************************************************
gogetrfc.go
--------------------------------------------------------------------------------
RFC (Request For Comment) Downloader.
Takes RFC #'s as positional args (numbers only, whitespace sep).
Prints rfc to stout (TODO: write to a file called rfc####.txt).
CAUTION: uses http cleartext communication.
http://brianc2788.github.io/
******************************************************************************
*/
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

const (
	RFC_APIURL_1 = "https://www.rfc-editor.org/rfc/"
	RFC_APIURL_2 = "" //tools.ietf.org ?
)

func main() {
	ArgList := ParseArgs()

	for _, url := range ArgList {
		content, err := getRfcText(url)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s", content)
	}
}

func ParseArgs() []string {
	var args []string
	if len(os.Args) > 1 {
		args = os.Args[1:]
	}
	return args
}

func getRfcText(rfcnum string) (string, error) {
	url := RFC_APIURL_1 + "rfc" + rfcnum + ".txt"

	fmt.Printf("URL: %s\n", url)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("status code: %d\n", resp.StatusCode)
		return "", err
	}

	ct := resp.Header.Get("Content-Type")
	if !strings.Contains(ct, "text/plain") {
		err := fmt.Errorf("%T\nRecieved incorrect content-type: %s\n", resp, ct)
		return "", err
	}

	txt, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	return string(txt), nil
}
