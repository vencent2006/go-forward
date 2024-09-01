import { Injectable, UnauthorizedException } from "@nestjs/common";
import { JwtService } from "@nestjs/jwt";
import { AccountService } from "../account/account.service";

@Injectable()
export class AuthService {
  constructor(
    private accountService: AccountService,
    private jwtService: JwtService) {
  }

  async login(username, password) {
    const user = await this.accountService.findOneByUsername(username);
    // const md5password = md5(password).toUpperCase();
    // console.log(user, md5password);
    if (!user || user.password !== password) {
      throw new UnauthorizedException();
    }
    const payload = { username: user.username, userid: user.id };
    return {
      token: await this.jwtService.signAsync(payload)
    };
  }
}
