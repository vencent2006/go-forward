import { Length } from "class-validator";


export class CreateAccountDto {

  @Length(5, 10, { message: "用户名长度在5~10个字符" })
  username: string;

  @Length(6, 20, { message: "密码长度在6~20个字符" })
  password: string;

  @Length(6, 20, { message: "昵称长度在6~20个字符" })
  nickname: string;

  @Length(6, 20, { message: "头像长度在6~20个字符" })
  avatar: string;
}
