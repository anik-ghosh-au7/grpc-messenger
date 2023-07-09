const grpc = require('@grpc/grpc-js'); // Import the gRPC JavaScript library
const protoLoader = require('@grpc/proto-loader'); // Import the gRPC Proto Loader
const readline = require('readline'); // Import Node's readline module

// Create a readline interface for reading from and writing to the standard input and output
const rl = readline.createInterface({
  input: process.stdin,
  output: process.stdout
});

// Ask the user for the server URL
rl.question("Enter the server url (example: localhost:8080): ", (serverURL) => {
  // Ask the user for their client ID
  rl.question("Enter your client ID: ", (clientID) => {
    // Load the chat.proto file
    const packageDefinition = protoLoader.loadSync('../proto/chat.proto');

    // Load the chat service definition from the package definition
    const { main } = grpc.loadPackageDefinition(packageDefinition);

    // Connect to the gRPC server using the chat service
    const client = new main.ChatApi(
      serverURL,
      grpc.credentials.createInsecure()
    );

    // Create a new stream to the Connect method on the server
    const stream = client.Connect({ id: clientID });

    // Listen for 'data' events on the stream, which are triggered when messages are received from the server
    stream.on('data', (message) => {
      // Log incoming messages
      console.log(`${message.user.id}: ${message.content}`);
    });

    // Listen for 'error' events on the stream, which are triggered when an error occurs
    stream.on('error', (error) => {
      // Log the error and exit the process
      console.log('Disconnected from server due to error:', error.message);
      process.exit(1);
    });

    // Listen for 'end' events on the stream, which are triggered when the server ends the stream
    stream.on('end', () => {
      // Log the disconnection and exit the process
      console.log('Disconnected from server.');
      process.exit();
    });

    // Listen for 'line' events on the readline interface, which are triggered when the user enters a message
    rl.on('line', (line) => {
      // Send the message to the server using the Broadcast method
      client.Broadcast(
        { user: { id: clientID }, content: line.trim() },
        (err) => {
          if (err) {
            // If there's an error, log it
            console.log('Error broadcasting message:', err);
          }
        }
      );
    });
  });
});
