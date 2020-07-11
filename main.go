package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
)

/*
csv file format:
short_postfix,url,public

public field - zero is private, non-zero is public

example:
wiki,https://bendun.cc/wiki,0
*/

const csvURLDatabaseFilename = "urls.csv"
const defaultURL = "https://bendun.cc/"

func main() {
	csvFilename := flag.String("csv-file", csvURLDatabaseFilename, "Source of URLs and their short identifiers.")
	port := flag.Int("port", 0, "Port to listen for requests")
	flag.Parse()

	db, err := readURLDatabase(*csvFilename)
	if err != nil {
		log.Fatal(err)
	}

	urlEntries = db
	http.HandleFunc("/", redirect)

	// get random availibe port from OS
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Listening at port: ", listener.Addr().(*net.TCPAddr).Port)
	go log.Fatal(http.Serve(listener, nil))
}
