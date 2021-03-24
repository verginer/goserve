// Miniserver is the simplest go server serving pages from a directory
package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	// create file server handler

	port := flag.String("p", "8888", "port to use")
	myDir := flag.String("d", ".", "directory to serve from")
	flag.Parse()

	fs := http.FileServer(http.Dir(*myDir))

	// start HTTP server with `fs` as default handler
	log.Print("Serving " + *myDir + " at http://localhost:" + *port)
	log.Fatal(http.ListenAndServe(":"+*port, fs))

}
