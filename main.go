package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func main() {
	resCode := http.StatusOK

	if len(os.Args) > 1 {
		code, err := strconv.Atoi(os.Args[1])
		if err == nil {
			resCode = code
		}
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("\nreceived a new request:\n")

		// Print the request URL, headers and body
		fmt.Printf("path: %s\n", r.URL.Path)

		fmt.Printf("headers:\n")
		for k, v := range r.Header {
			fmt.Printf("  %s: %s\n", k, strings.Join(v, ", "))
		}

		fmt.Printf("body:\n")
		io.Copy(os.Stdout, r.Body)

		fmt.Printf("\n")

		w.WriteHeader(resCode)
	})

	log.Fatal(http.ListenAndServe(":3001", nil))
}
