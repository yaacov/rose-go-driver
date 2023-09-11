package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/RedHat-Israel/rose-go-driver/pkg/server"
	"github.com/sirupsen/logrus"
)

// main is the entry point of the application. It retrieves the server address and logging level
// from command line flags, sets the logging level, and starts the server.
func main() {
	addr, logLevel := parseCmdFlags()
	setLoggingLevel(logLevel)
	startServer(addr)
}

// parseCmdFlags parses the command line flags to determine various configuration parameters,
// including the server's address and the desired logging level. It uses the "port", "listen",
// and "loglevel" flags.
//
// Returns:
//   - string: The server address in the format "listen:port".
//   - string: The desired logging level as a string.
func parseCmdFlags() (string, string) {
	// Define command line arguments
	port := flag.Int("port", 8081, "Specify the port number.")
	listen := flag.String("listen", "", "Specify the listen address. Default is all interfaces.")
	logLevel := flag.String("loglevel", "info", "Set the logging level. Options: debug, info, warn, error, fatal, panic")

	flag.Parse()

	addr := fmt.Sprintf("%s:%d", *listen, *port)
	return addr, *logLevel
}

// setLoggingLevel sets the logging level for the application based on the provided logLevel string.
//
// Parameters:
//   - logLevel: The desired logging level as a string.
func setLoggingLevel(logLevel string) {
	level, err := logrus.ParseLevel(logLevel)
	if err != nil {
		logrus.Fatalf("Invalid log level: %v", err)
	}

	logrus.SetLevel(level)
	logrus.Infof("Logging level set to %s", level)
}

// startServer initializes the HTTP server and starts listening on the provided address.
// It also sets up the necessary routes for handling incoming requests.
//
// Parameters:
//   - addr: The address on which the server should listen.
func startServer(addr string) {
	log.Printf("Starting web server on %s", addr)
	log.Printf("Access the server at http://127.0.0.1:%s", strings.Split(addr, ":")[1])

	http.HandleFunc("/", server.HandleRequest)
	log.Fatal(http.ListenAndServe(addr, nil))
}
