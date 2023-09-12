// Package main is the entry point for the GFW Knock application.
package main

import (
	"fmt"
	"gfwKnock/pkg/config"
	"gfwKnock/pkg/connection"
	"gfwKnock/pkg/ulimit"
	"log"
)

var (
	codename = "GFW Knock , GFW Knock Bypass PoC."
)

func main() {
	// print name & code name
	fmt.Println("\u001B[92m", codename, "[0m")

	config.Load(".")
	server := connection.Server{}

	server = server.InitialConfiguration()

	err := ulimit.SetULimit()
	if err != nil {
		log.Fatalf("error : %s\n", err.Error())
	}

	server.Bind(fmt.Sprintf("127.0.0.1:%d", server.ListenPort))
}
