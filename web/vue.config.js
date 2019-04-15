const CompressionPlugin = require("compression-webpack-plugin");

module.exports = {
  pages: {
    index: {
      entry: "src/entries/index.js",
      template: 'public/index.html',
      filename: 'index.html',
      title: 'Index Page',
      chunks: ['chunk-vendors', 'chunk-common', 'index']
    },
    initialize: {
      entry: "src/entries/initialize.js",
      template: 'public/index.html',
      filename: 'initialize.html',
      title: 'Insatll Page',
      chunks: ['chunk-vendors', 'chunk-common', 'initialize']
    }
  },
  devServer: {
    port: 8080,
    proxy: {
      '/api': {
        target: 'http://localhost:9701/',
        ws: true,
        changeOrigin: true,
        pathRewrite: {
          '^/api': '/api'
        }
      }
    }
  },
  // configureWebpack: config => {
  //   if (process.env.NODE_ENV === 'production') {
  //     config.plugins.push(new CompressionPlugin({
  //         algorithm: 'gzip',
  //         test: /\.js$|\.css/,
  //         threshold: 10240,
  //         minRatio: 0.8,
  //         deleteOriginalAssets: true
  //       })
  //     )
  //   }
  // }
};
