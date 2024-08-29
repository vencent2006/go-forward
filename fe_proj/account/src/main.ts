import { NestFactory } from "@nestjs/core";
import { AppModule } from "./app.module";
import { ValidationPipe } from "@nestjs/common";
import { Response } from "./common/response";

async function bootstrap() {
  const app = await NestFactory.create(AppModule);
  app.useGlobalInterceptors(new Response()) // 全局响应拦截器
  app.useGlobalPipes(new ValidationPipe());// 注册全局DTO验证通道
  await app.listen(3000);
}

bootstrap();
