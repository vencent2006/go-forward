import { Injectable, UnauthorizedException } from "@nestjs/common";
import { UserService } from "../user/user.service";
import * as md5 from "md5"
import { JwtService } from "@nestjs/jwt"; // 这个要这么写

@Injectable()
export class AuthService {
  constructor(
    private userService: UserService,
    private jwtService: JwtService,) {
  }
  async login(username, password){
    const user = await this.userService.findByUsername(username)
    const md5password = md5(password).toUpperCase()
    console.log(user, md5password);
    if(!user || user.password !== md5password){
      throw new UnauthorizedException();
    }
    const payload = {username: user.username, userid: user.id};
    return {
      token: await this.jwtService.signAsync(payload)
    }
  }
}
