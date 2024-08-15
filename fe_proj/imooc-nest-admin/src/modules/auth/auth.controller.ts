import { Controller, Post } from "@nestjs/common";
import { Public } from "./public.decorator";

@Controller('auth')
export class AuthController {
  @Public()
  @Post('login')
  login(){
    return 'do login';
  }
}
