package server

import (
	"context"
	"log"
	"net"
	"sync"

	"github.com/anik-ghosh-au7/grpc-messenger/gen/chat"
	"google.golang.org/grpc"
)

type server struct {
	chat.UnimplementedChatApiServer
	clients map[string]chat.ChatApi_ConnectServer
	mu      sync.Mutex
}

func (s *server) Connect(user *chat.User, stream chat.ChatApi_ConnectServer) error {
	s.mu.Lock()
	s.clients[user.Id] = stream
	s.mu.Unlock()
	log.Println("Client Connected: ", user.Id)
	<-stream.Context().Done()
	s.mu.Lock()
	delete(s.clients, user.Id)
	s.mu.Unlock()
	log.Println("Client Disconnected: ", user.Id)
	return nil
}

func (s *server) Broadcast(ctx context.Context, message *chat.Message) (*chat.Message, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	for id, clientStream := range s.clients {
		if id != message.User.Id {
			if err := clientStream.Send(message); err != nil {
				log.Println("Error broadcasting message to", id, ":", err)
			}
		}
	}
	return message, nil
}

func StartGrpcServer() error {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		return err
	}

	srv := grpc.NewServer()
	chat.RegisterChatApiServer(srv, &server{
		clients: make(map[string]chat.ChatApi_ConnectServer),
	})

	log.Println("Server started. Listening on port 8080.")
	return srv.Serve(listener)
}
