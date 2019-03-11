module.exports = {
  pages: {
    index: {
      entry: "src/entries/index.js",
      template: 'public/index.html',
      filename: 'index.html',
      title: 'Index Page',
      chunks: ['chunk-vendors', 'chunk-common', 'index']
    },
    install: {
      entry: "src/entries/install.js",
      template: 'public/index.html',
      filename: 'install.html',
      title: 'Insatll Page',
      chunks: ['chunk-vendors', 'chunk-common', 'install']
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
  }
};
