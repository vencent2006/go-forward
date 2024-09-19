import {  NestFactory } from "@nestjs/core";
import { AppModule } from './app.module';
import { ProxyMiddleware } from "./proxy.middleware";

async function bootstrap() {
  const app = await NestFactory.create(AppModule);
  // app.use(ProxyMiddleware)
  await app.listen(3000);
}
bootstrap();
