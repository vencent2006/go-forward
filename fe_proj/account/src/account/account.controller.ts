import { Body, Controller, Get, HttpException, HttpStatus, Param, Post } from "@nestjs/common";
import { AccountService } from "./account.service";
import { CreateAccountDto } from "./dto/create-account.dto";
import { LoginAccountDto } from "./dto/login-account.dto";
import { AccountEntity } from "./entities/account.entity";
import { Public } from "../auth/public.decorator";

@Controller("account")
export class AccountController {
  constructor(private readonly accountService: AccountService) {
  }

  /**
   * 注册
   * @param createAccountDto
   */
  @Public()
  @Post("register")
  async register(@Body() createAccountDto: CreateAccountDto) {
    const accountExisted: AccountEntity = await this.accountService.findOneByUsername(createAccountDto.username);
    if (accountExisted) {
      throw new HttpException("username(" + createAccountDto.username + ") already existed", HttpStatus.BAD_REQUEST);
    }
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
