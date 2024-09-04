import { Body, Controller, Get, Post, Query } from "@nestjs/common";
import { AuthService } from "./auth.service";
import { Public } from "./public.decorator";
import { LoginAccountDto } from "../account/dto/login-account.dto";

@Controller("auth")
export class AuthController {
  constructor(private readonly authService: AuthService) {
  }

  @Public()
  @Post("login")
  login(@Body() loginAccountDto: LoginAccountDto) {
    return this.authService.login(loginAccountDto);
  }

  /**
   * 刷新
   * @param refreshToken
   */
  @Get("refresh")
  refreshToken(@Query("refresh_token") refreshToken: string) {
    return this.authService.refreshToken(refreshToken);
  }
}
