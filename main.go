package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/anik-ghosh-au7/grpc-messenger/client"
	"github.com/anik-ghosh-au7/grpc-messenger/server"
)

func main() {
	reader := bufio.NewReader(os.Stdin) // Create a new reader to read from the standard input
	// Ask the user whether to start the application as a server or as a client
	fmt.Print("Enter 's' to start as a server, or 'c' to start as a client: ")
	option, _ := reader.ReadString('\n') // Read the user's option from the standard input
	option = strings.TrimSpace(option)   // Trim whitespace from the option

	// Depending on the user's choice, start the application as a server or as a client
	switch option {
	case "s":
		// Start the gRPC server
		err := server.StartGrpcServer()
		if err != nil {
			log.Fatalf("Failed to start the server: %v", err) // If there's an error, log it and exit
		}
	case "c":
		// Start the gRPC client
		err := client.StartGrpcClient()
		if err != nil {
			log.Fatalf("Failed to start the client: %v", err) // If there's an error, log it and exit
		}
	default:
		fmt.Println("Invalid option. Exiting.") // If the user's option is not "s" or "c", print an error message and exit
	}
}
