import { Module } from '@nestjs/common';
import { AuthController } from './auth.controller';
import { AuthService } from "./auth.service";
import { APP_GUARD } from "@nestjs/core";
import { AuthGuard } from "./auth.guard";
import { UserModule } from "../user/user.module";

@Module({
  imports: [UserModule], // 引用了module
  controllers: [AuthController],
  providers:[
    AuthService,
    {
      provide: APP_GUARD,
      useClass: AuthGuard, // 守卫 guard
    }]
})
export class AuthModule {}
