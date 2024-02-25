package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	"os"
)

var addr = flag.String("addr", "google.com", "address to resolve")

func init() {
	_ = os.Setenv("GODEBUG", "netdns=2")
}

func main() {
	flag.Parse()

	addrs, err := net.DefaultResolver.LookupIPAddr(context.Background(), *addr)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, addr := range addrs {
		fmt.Printf("- %s\n", addr.String())
	}
}
