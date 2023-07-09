const path = require('path');
const nodeExternals = require('webpack-node-externals');

module.exports = {
  target: 'node',
  mode: 'production',
  externals: [nodeExternals()],
  entry: {
    app: ['./client.js']
  },
  output: {
    path: path.resolve(__dirname, 'dist'),
    filename: 'client.bundle.js'
  }
};
