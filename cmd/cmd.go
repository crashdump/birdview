package main

import (
	"flag"
	"fmt"
	"github.com/crashdump/birdview/cmd/api/v1"
	"log"
	"net/http"
	"os"
	"time"
)

var usage = `usage: birdview [--version] [--help]`

func main() {
	help := flag.Bool("help", false, "display the usage")
	version := flag.Bool("version", false, "display the version")
	flag.Parse()

	if *help {
		fmt.Println(usage)
		os.Exit(0)
	}

	if *version {
		fmt.Println("v1.0.0")
		os.Exit(0)
	}

	srv := &http.Server{
		Addr:         "127.0.0.1:8080",
		Handler:      v1.Router(),
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	log.Println("Routes parsed, webserver starting on :8080")
	log.Fatal(srv.ListenAndServe())
}
