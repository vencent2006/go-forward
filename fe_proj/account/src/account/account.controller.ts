import { Controller, Get, Post, Body, Param } from "@nestjs/common";
import { AccountService } from "./account.service";
import { CreateAccountDto } from "./dto/create-account.dto";
import { LoginAccountDto } from "./dto/login-account.dto";

@Controller("account")
export class AccountController {
  constructor(private readonly accountService: AccountService) {
  }

  /**
   * 注册
   * @param createAccountDto
   */
  @Post("register")
  register(@Body() createAccountDto: CreateAccountDto) {
    return this.accountService.create(createAccountDto);
  }

  /**
   * 登录
   * @param loginAccountDto
   */
  @Post("login")
  login(@Body() loginAccountDto: LoginAccountDto) {
    return this.accountService.login(loginAccountDto);
  }


  /**
   * 获取全部账号
   */
  @Get()
  findAll() {
    return this.accountService.findAll();
  }

  /**
   * 获取某一个账号
   * @param id
   */
  @Get(":id")
  findOne(@Param("id") id: string) {
    return this.accountService.findOne(+id);
  }

}
