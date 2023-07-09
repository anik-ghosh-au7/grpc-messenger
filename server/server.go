package server

import (
	"context"
	"log"
	"net"
	"net/http"
	"sync"

	"github.com/anik-ghosh-au7/grpc-messenger/gen/chat"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

// Define a server struct which includes a mutex and a map of clients.
type server struct {
	chat.UnimplementedChatApiServer                                       // Embed the gRPC server interface
	clients                         map[string]chat.ChatApi_ConnectServer // Map of clients, keys are user IDs, values are stream interfaces
	mu                              sync.Mutex                            // Mutex for concurrent operations
}

// Connect is a method on the server struct which sets up a new client and watches for disconnection.
func (s *server) Connect(user *chat.User, stream chat.ChatApi_ConnectServer) error {
	s.mu.Lock()                                   // Lock the mutex to prevent concurrent map writes
	s.clients[user.Id] = stream                   // Add the client to the map
	s.mu.Unlock()                                 // Unlock the mutex
	log.Println("Client Connected: ", user.Id)    // Log the connection
	<-stream.Context().Done()                     // Wait for the client to be done
	s.mu.Lock()                                   // Lock the mutex
	delete(s.clients, user.Id)                    // Remove the client from the map
	s.mu.Unlock()                                 // Unlock the mutex
	log.Println("Client Disconnected: ", user.Id) // Log the disconnection
	return nil                                    // Return nil since there's no error to return
}

// Broadcast is a method on the server struct which broadcasts a message to all connected clients.
func (s *server) Broadcast(ctx context.Context, message *chat.Message) (*chat.Message, error) {
	s.mu.Lock()         // Lock the mutex to prevent concurrent map reads
	defer s.mu.Unlock() // Defer unlocking the mutex
	// Loop over all clients and send the message to each client (except for the sender)
	for id, clientStream := range s.clients {
		if id != message.User.Id {
			if err := clientStream.Send(message); err != nil {
				log.Println("Error broadcasting message to", id, ":", err) // Log the error if there's one
			}
		}
	}
	return message, nil // Return the message and nil error
}

// GetClients returns the IDs of all connected clients.
func (s *server) GetClients(ctx context.Context, empty *chat.Empty) (*chat.ClientList, error) {
	s.mu.Lock() // Prevent concurrent map reads
	defer s.mu.Unlock()
	clientIDs := make([]string, 0, len(s.clients))
	for id := range s.clients {
		clientIDs = append(clientIDs, id)
	}
	return &chat.ClientList{
		ClientIds: clientIDs,
	}, nil
}

// StartGrpcServer is a function which starts a new gRPC server.
func StartGrpcServer() error {
	s := &server{ // Create a new server instance
		clients: make(map[string]chat.ChatApi_ConnectServer), // Initialize an empty clients map
	}
	go func() {
		mux := runtime.NewServeMux()                                    // Create a new ServeMux
		chat.RegisterChatApiHandlerServer(context.Background(), mux, s) // Register the server on the mux
		log.Fatalln(http.ListenAndServe(":8000", mux))                  // Start the HTTP server
	}()
	listener, err := net.Listen("tcp", ":8080") // Listen for TCP connections on port 8080
	if err != nil {
		return err // If there's an error, return it
	}
	srv := grpc.NewServer() // Create a new gRPC server
	chat.RegisterChatApiServer(srv, s)
	log.Println("Server started. Listening on port 8080.") // Log server start
	return srv.Serve(listener)                             // Start serving the server and return any potential error
}
