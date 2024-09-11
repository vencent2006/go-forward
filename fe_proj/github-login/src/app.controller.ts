import { Controller, Get, Req, UseGuards } from "@nestjs/common";
import { AppService } from './app.service';
import { AuthGuard } from "@nestjs/passport";

@Controller()
export class AppController {
  constructor(private readonly appService: AppService) {}

  @Get('login')
  @UseGuards(AuthGuard('github'))
  async login(){
    return
  }

  @Get("callback")
  @UseGuards(AuthGuard('github'))
  async authCallback(@Req() req) {
    console.log(req);
    return req.user
  }
}
