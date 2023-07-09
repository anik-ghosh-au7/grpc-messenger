package client

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/anik-ghosh-au7/grpc-messenger/gen/chat"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// StartGrpcClient is a function that starts a new gRPC client.
func StartGrpcClient() error {
	reader := bufio.NewReader(os.Stdin)                           // Create a new reader to read from the standard input
	fmt.Print("Enter the server url (example: localhost:8080): ") // Ask the user to enter the server URL
	serverURL, _ := reader.ReadString('\n')                       // Read the server URL from the standard input
	serverURL = strings.TrimSpace(serverURL)                      // Trim whitespace from the server URL
	// Dial a connection to the gRPC server with insecure transport credentials
	conn, err := grpc.Dial(serverURL, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err // If there's an error, return it
	}
	defer conn.Close()                     // Ensure the connection will be closed when the function exits
	client := chat.NewChatApiClient(conn)  // Create a new chat client
	fmt.Print("Enter your client ID: ")    // Ask the user to enter their client ID
	clientID, _ := reader.ReadString('\n') // Read the client ID from the standard input
	clientID = strings.TrimSpace(clientID) // Trim whitespace from the client ID
	user := &chat.User{Id: clientID}       // Create a new user with the client ID
	// Open a new stream with the server. This stream will be used to send and receive messages.
	stream, err := client.Connect(context.Background(), user)
	if err != nil {
		return err // If there's an error, return it
	}
	// Start a new goroutine to handle incoming messages
	go func() {
		for {
			// Receive a message from the server
			message, err := stream.Recv()
			if err != nil {
				fmt.Println("Disconnected from server.") // If there's an error, the client has disconnected
				return                                   // Exit the goroutine
			}
			// Print the message content
			fmt.Println(message.User.Id+": ", message.Content)
		}
	}()
	// Start an infinite loop to handle outgoing messages
	for {
		// Read a message from the standard input
		messageContent, _ := reader.ReadString('\n')
		// Create a new message
		msg := &chat.Message{
			User:    user,
			Content: strings.TrimSpace(messageContent), // Trim whitespace from the message content
		}
		// Send the message to the server
		client.Broadcast(context.Background(), msg)
	}

}
