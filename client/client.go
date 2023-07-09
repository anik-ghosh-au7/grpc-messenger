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

func StartGrpcClient() error {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the server url (example: localhost:8080): ")
	serverURL, _ := reader.ReadString('\n')
	serverURL = strings.TrimSpace(serverURL)

	conn, err := grpc.Dial(serverURL, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}
	defer conn.Close()

	client := chat.NewChatApiClient(conn)

	fmt.Print("Enter your client ID: ")
	clientID, _ := reader.ReadString('\n')
	clientID = strings.TrimSpace(clientID)

	user := &chat.User{Id: clientID}
	stream, err := client.Connect(context.Background(), user)
	if err != nil {
		return err
	}

	go func() {
		for {
			message, err := stream.Recv()
			if err != nil {
				fmt.Println("Disconnected from server.")
				return
			}
			fmt.Println(message.User.Id+": ", message.Content)
		}
	}()

	for {
		messageContent, _ := reader.ReadString('\n')
		msg := &chat.Message{
			User:    user,
			Content: strings.TrimSpace(messageContent),
		}
		client.Broadcast(context.Background(), msg)
	}

}
