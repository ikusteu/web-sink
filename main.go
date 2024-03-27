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
	// Initial value for response code for every request: 200
	resCode := http.StatusOK

	// If a response code is passed as an argument, use it
	if len(os.Args) > 1 {
		code, err := strconv.Atoi(os.Args[1])
		if err == nil {
			resCode = code
		}
	}

	// Create server instance
	server := &Server{resCode}

	// Start a server in a goroutine (akin to offloading it to a separate thread)
	go server.Start()

	// Start CLI for user input - this IS on the main thread, is listening for user input,
	// for commands like: 'setres <response_code>' and 'clear'
	//
	// This is blocking (infinite loop):
	// - listening to user input
	// - preventing the program from exiting until SIGINT, SIGKILL and such
	StartCLI(server)
}

// A struct used to represent the server:
// - it holds response code (which we can easily change between requests, without killing the server)
// - it implements http.Handler interface (has .serveHTTP method) so that we can use it to process requests received by default server mux (listening to all paths)
type Server struct {
	resCode int
}

// A function used to handle the request
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// fmt.Printf is akin to printf of stdio.h in C
	fmt.Printf("\nreceived a new request:\n")

	// Print the request URL, headers and body
	fmt.Printf("path: %s\n", r.URL.Path)

	fmt.Printf("headers:\n")
	for k, v := range r.Header {
		fmt.Printf("  %s: %s\n", k, strings.Join(v, ", "))
	}

	fmt.Printf("body:\n")

	// io.Copy is used to copy from a reader to a writer (in this case it reads from the request body and writes to stdout)
	io.Copy(os.Stdout, r.Body)

	fmt.Printf("\n")

	// Return the resonse and close the connection
	w.WriteHeader(s.resCode)
}

// A simple wrapper attaching the server to port :3001 of the defaul server mux
func (s *Server) Start() {
	log.Fatal(http.ListenAndServe(":3001", s))
}

// A helper used to set the response code for all subsequent requests
func (s *Server) SetResponse(code int) {
	s.resCode = code
}

// Starts a CLI for user input: handles commands:
// - setres <response_code> - sets the response code for all subsequent requests
// - clear - clears the console (this is helpful between requests)
func StartCLI(server *Server) {
	// Akin to 'while (true)'
	for {
		// Buffer to read the command from stdin into
		buf := make([]byte, 256)
		// Read the command
		n, err := os.Stdin.Read(buf)
		if err != nil {
			log.Fatal(err)
		}

		// The rest is self-explanatory

		input := strings.Split(strings.TrimSpace(string(buf[:n])), " ")
		cmd := input[0]

		switch cmd {
		case "setres":
			value := input[1]
			code, err := strconv.Atoi(value)
			if err != nil {
				fmt.Printf("Invalid response code\n")
				continue
			}
			server.SetResponse(code)
		case "clear":
			os.Stdout.WriteString("\033[H\033[2J")
		default:
			fmt.Printf("Invalid command\n")
		}
	}
}
