import { Module } from '@nestjs/common';
import { AuthController } from './auth.controller';
import { AuthService } from "./auth.service";
import { APP_GUARD } from "@nestjs/core";
import { AuthGuard } from "./auth.guard";
import { UserModule } from "../user/user.module";
import { JwtModule } from "@nestjs/jwt";
import { JWT_SECRET_KEY } from "./auth.jwt.secret";

@Module({
  imports: [
    UserModule,
    JwtModule.register(
    {
      global: true,
      secret: JWT_SECRET_KEY, // 私钥
      signOptions: {expiresIn: 24 * 60 * 60 + 's'} // 一天过期
    }
  )], // 引用了module
  controllers: [AuthController],
  providers:[
    AuthService,
    {
      provide: APP_GUARD,
      useClass: AuthGuard, // 守卫 guard
    }]
})
export class AuthModule {}
