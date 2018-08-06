const path = require('path')
const HtmlWebPackPlugin = require("html-webpack-plugin")

const htmlPlugin = new HtmlWebPackPlugin({
  template: "./src/index.html",
  filename: "./index.html"
})

module.exports = {
  entry: './src/application.js',
  mode: 'development',
  devtool: 'source-map',
  watch: true,
  output: {
    filename: 'main.js',
    path: path.resolve(__dirname, 'dist')
  },
  module: {
    rules: [
      {
        test: /\.js$|\.jsx/,
        exclude: /node_modules/,
        use: {
          loader: "babel-loader"
        }
      }
    ]
  },
  plugins: [htmlPlugin]
}
