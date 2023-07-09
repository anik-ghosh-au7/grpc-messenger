const grpc = require('@grpc/grpc-js');
const protoLoader = require('@grpc/proto-loader');
const readline = require('readline');

const rl = readline.createInterface({
  input: process.stdin,
  output: process.stdout
});

// Ask for server URL
rl.question("Enter the server url (example: localhost:8080): ", (serverURL) => {
  // Ask for client ID
  rl.question("Enter your client ID: ", (clientID) => {
    // Load chat.proto file
    const packageDefinition = protoLoader.loadSync('../gen/chat/chat.proto', {
      keepCase: true,
      longs: String,
      enums: String,
      defaults: true,
      oneofs: true
    });

    const chat = grpc.loadPackageDefinition(packageDefinition);

    // Connect to the server
    const client = new chat.ChatApi(
      serverURL,
      grpc.credentials.createInsecure()
    );

    // Create the Connect stream
    const stream = client.Connect({ id: clientID });

    stream.on('data', (message) => {
      // Log incoming messages
      console.log(`${message.user.id}: ${message.content}`);
    });

    stream.on('error', (error) => {
      // Log if the server is stopped or any other error occurs
      console.log('Disconnected from server due to error:', error.message);
      process.exit(1);
    });

    stream.on('end', () => {
      // Log if the server ends the stream
      console.log('Disconnected from server.');
      process.exit();
    });

    rl.on('line', (line) => {
      // Send messages to the server
      client.Broadcast(
        { user: { id: clientID }, content: line.trim() },
        (err) => {
          if (err) {
            console.log('Error broadcasting message:', err);
          }
        }
      );
    });
  });
});
