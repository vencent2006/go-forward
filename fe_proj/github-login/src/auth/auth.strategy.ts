import { Injectable } from "@nestjs/common";
import { PassportStrategy } from "@nestjs/passport";
import { Profile, Strategy } from "passport-github2";


@Injectable()
export class GithubStrategy extends PassportStrategy(Strategy, 'github') {
  constructor() {
    super({
      clientID: 'Ov23liDY828hnLDm46Pq',
      clientSecret: '0d5afa5a7f862353c756261746a8f601834199a3',
      callbackURL: 'http://localhost:3000/callback',
      scope: ['public_profile'], // 作用域
    });
  }

  async validate(accessToken: string, refreshToken: string, profile: Profile) {
    console.log(profile, accessToken, refreshToken);
    return profile
  }
}
