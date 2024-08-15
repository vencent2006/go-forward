import { Injectable, UnauthorizedException } from "@nestjs/common";
import { UserService } from "../user/user.service";

@Injectable()
export class AuthService {
  constructor(private userService: UserService) {
  }
  async login(username, password){
    const user = await this.userService.findByUsername(username)
    console.log(user);
  }
}
