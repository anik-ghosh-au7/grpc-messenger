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
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter 's' to start as a server, or 'c' to start as a client: ")
	option, _ := reader.ReadString('\n')
	option = strings.TrimSpace(option)

	switch option {
	case "s":
		err := server.StartGrpcServer()
		if err != nil {
			log.Fatalf("Failed to start the server: %v", err)
		}
	case "c":
		err := client.StartGrpcClient()
		if err != nil {
			log.Fatalf("Failed to start the client: %v", err)
		}
	default:
		fmt.Println("Invalid option. Exiting.")
	}
}
