package main

import (
	"fmt"
	"net"
	"os"
	"time"

	"golang.org/x/net/context"
)

func main() {
	// Set a timeout for the DNS lookup.
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Perform a DNS lookup for the IP address of a domain.
	ipAddr, err := net.DefaultResolver.LookupIPAddr(ctx, "google.com")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err.Error())
		os.Exit(1)
	}

	// Print the first IP address returned by the DNS lookup.
	fmt.Printf("IP address for google.com: %s\n", ipAddr[0].IP.String())
}
