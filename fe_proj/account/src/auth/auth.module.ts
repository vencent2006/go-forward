import { Module } from "@nestjs/common";
import { AuthService } from "./auth.service";
import { JwtModule } from "@nestjs/jwt";
import { AccountModule } from "../account/account.module";
import { JWT_SECRET_KEY, JWT_ACCESS_TOKEN_PERIOD } from "./auth.jwt.secret";
import { APP_GUARD } from "@nestjs/core";
import { AuthController } from "./auth.controller";
import { AuthGuard } from "./auth.guard";

@Module({
  imports: [
    AccountModule,
    JwtModule.register(
      {
        global: true,
        secret: JWT_SECRET_KEY, // 私钥
        signOptions: { expiresIn: JWT_ACCESS_TOKEN_PERIOD } // 过期时间
      }
    )], // 引用了module
  controllers: [AuthController],
  providers: [
    AuthService,
    {
      provide: APP_GUARD,
      useClass: AuthGuard // 守卫 guard
    }]
})
export class AuthModule {
}
