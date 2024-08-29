import { Length } from "class-validator";

export class LoginAccountDto {
  @Length(5, 10, { message: "用户名长度在5~10个字符" })
  username: string;

  @Length(6, 10, { message: "密码长度在6~10个字符" })
  password: string;
}
