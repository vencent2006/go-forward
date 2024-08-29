import { NestFactory } from "@nestjs/core";
import { AppModule } from "./app.module";
import { Response } from "./common/response";
import { ValidationPipe } from "@nestjs/common";
import { HttpFilter } from "./common/filter";

async function bootstrap() {
  const app = await NestFactory.create(AppModule);
  app.useGlobalFilters(new HttpFilter()) // 全局异常拦截器
  app.useGlobalInterceptors(new Response()) // 全局响应拦截器
  app.useGlobalPipes(new ValidationPipe());// 注册全局DTO验证通道
  await app.listen(3000);
}

bootstrap();
