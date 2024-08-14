import { Controller, Get } from "@nestjs/common";

@Controller('auth')
export class AuthController {
  @Get()
  auth(){
    return 'auth';
  }
}
