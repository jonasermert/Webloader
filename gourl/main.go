package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
)

var (
	flagOutput = flag.String("o", "", "file as target")
	flagHeader = flag.Bool("header", false, "output of the http-header")
)

func main() {

	flag.Parse()
	args := flag.Args()
	var w io.Writer
	// default is stdout
	w = os.Stdout
	if len(args) != 1 {
		fmt.Println("Please enter only one URL as input")
		os.Exit(1)
	}
	// fmt.Println(*flagOutput, os.Args, flag.Args())

	url := args[0]
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("error loading data", url, err)
		os.Exit(2)
	}
	defer resp.Body.Close()
	if *flagOutput != "" {
		path := filepath.Dir(*flagOutput)
		err := os.MkdirAll(path, 0755)
		if err != nil {
			fmt.Println("cannot create folder", path, err)
			os.Exit(22)
		}
		fd, err := os.OpenFile(*flagOutput, os.O_CREATE|os.O_WRONLY, 0755)
		if err != nil {
			fmt.Println("error creating file", err, *flagOutput)
			os.Exit(20)
		}
		defer fd.Close()
		w = fd
	}

	if *flagHeader {
		fmt.Printf("%#v", resp.Header)
		for k, v := range resp.Header {
			fmt.Printf("%s: %v\n\n", k, v)
		}
	}

	io.Copy(w, resp.Body)

}

func validateURL(s string) bool {
	_, err := ParseRequestURI(s)
	if err != nil {
		return false
	}
	return true
	validURL := regexp.MustCompile("^http(s)?://[[:graph:]]+")
	return validURL.MatchString(s)
	// return true
}
