// proxy.middleware.ts
import { RequestHandler } from 'express';
import { createProxyMiddleware } from 'http-proxy-middleware';

const proxyOptions = {
  target: 'http://127.0.0.1:7890', // 代理服务器的URL
  changeOrigin: true,
  // 其他代理选项...
};

const proxyMiddleware = createProxyMiddleware(proxyOptions);

export const ProxyMiddleware: RequestHandler = (req, res, next) => {
  return proxyMiddleware(req, res, next);
};
