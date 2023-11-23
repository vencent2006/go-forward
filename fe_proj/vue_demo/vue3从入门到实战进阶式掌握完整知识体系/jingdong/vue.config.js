const { defineConfig } = require('@vue/cli-service')
module.exports = defineConfig({
  transpileDependencies: true,
  // 这个publicPath 要具体看你前端代码访问的路径是什么
  // 比如前端代码在域名访问的时候，还需要加个jingdong，那么publicPath:'/jingdong'
  // 这个是配置静态资源访问的路径的，当然这些可以nginx的conf进行配置，让前端配置尽可能简单
  publicPath: './'
})
