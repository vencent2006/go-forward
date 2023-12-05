const { defineConfig } = require('@vue/cli-service')
const path = require('path')
function resolve(dir) {
  return path.join(__dirname, dir)
}
module.exports = defineConfig({
  // webpack devServer 提供了代理的功能，该 代理可以把所有请求到当前服务中的请求
  // 转发（代理）到另外的一个服务器上
  devServer: {
    proxy: {
      // 当地址中包含 /api 的时候，触发此代理
      // 用我自己的代理，自己写的server
      '/api': {
        // target: 'https://api.imooc-admin.lgdsunday.club/',
        target: 'http://localhost:8888/',
        changeOrigin: true
      }
    }
  },
  configureWebpack: {
    resolve: {
      fallback: {
        path: require.resolve('path-browserify')
      }
    }
  },
  chainWebpack(config) {
    // 首先把原来svg应用的loader(file-loader)先排除
    config.module.rule('svg').exclude.add(resolve('src/icons')).end()
    // 专门增加一个icons的规则
    config.module
      .rule('icons')
      .test(/\.svg$/)
      .include.add(resolve('src/icons'))
      .end()
      .use('svg-sprite-loader')
      .loader('svg-sprite-loader')
      .options({
        symbolId: 'icon-[name]'
      })
      .end()
  },
  transpileDependencies: true
})
