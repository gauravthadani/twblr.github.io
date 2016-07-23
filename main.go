package main

import (
	"flag"
	"fmt"
	"net/http"
)

func main() {

	flag.Usage = usage
	var port = flag.String("port", "1001", "port to run the file server on ")
	flag.Parse()

	s := fmt.Sprintf(":%s", port)
	panic(http.ListenAndServe(s, http.FileServer(http.Dir("./"))))
}

func usage() {
	fmt.Print(usagePrefix)
	flag.PrintDefaults()
}

var usagePrefix = `
File Serve
Usage:
package [options] <subcommand>
Options:
`
