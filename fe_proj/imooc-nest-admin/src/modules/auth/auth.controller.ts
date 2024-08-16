import { Body, Controller, Post, UseFilters } from "@nestjs/common";
import { Public } from "./public.decorator";
import { AuthService } from "./auth.service";
import { HttpExceptionFilter } from "../../exception/http-exception.filter";
import { error, success } from "../../utils";

@Controller('auth')
export class AuthController {
  constructor(private authService: AuthService) {}
  @Public()
  @Post('login')
  @UseFilters(new HttpExceptionFilter())
  login(@Body() params){
    // 这是个promise
    return this.authService.login(params.username, params.password)
      .then((data:any)=>success(data, '登录成功'))
      .catch((err:any)=>error(err.message))
  }
}
