import { HttpException, HttpStatus, Injectable } from "@nestjs/common";
import { CreateAccountDto } from "./dto/create-account.dto";
import { UpdateAccountDto } from "./dto/update-account.dto";
import { LoginAccountDto } from "./dto/login-account.dto";
import { InjectRepository } from "@nestjs/typeorm";
import { Repository } from "typeorm";
import { AccountEntity, AccountRole, AccountStatus } from "./entities/account.entity";

@Injectable()
export class AccountService {
  constructor(
    @InjectRepository(AccountEntity) // 注入repository
    private readonly accountRepository: Repository<AccountEntity>
  ) {
  }

  create(createAccountDto: CreateAccountDto) {
    const account = new AccountEntity();
    account.username = createAccountDto.username;
    account.password = createAccountDto.password;
    account.role = AccountStatus.VALID.valueOf();
    account.avatar = createAccountDto.avatar;
    account.nickname = createAccountDto.nickname;
    account.status = AccountRole.USER.valueOf();
    return this.accountRepository.save(account);
  }

  async login(loginAccountDto: LoginAccountDto) {
    const account = await this.findOneByUsername(loginAccountDto.username)
    // todo use md5 encrypted password
    if (!account || account.password != loginAccountDto.password){
      throw new HttpException("invalid login account", HttpStatus.BAD_REQUEST);
    }
    return "This action logins a new account";
  }

  findAll(): Promise<AccountEntity[]> {
    return this.accountRepository.find();
  }

  findOne(id: number): Promise<AccountEntity> {
    return this.accountRepository.findOneBy({ id });
  }

  findOneByUsername(username: string): Promise<AccountEntity> {
    return this.accountRepository.findOneBy({ username });
  }

  update(id: number, updateAccountDto: UpdateAccountDto) {
    return `This action updates a #${id} account`;
  }

  remove(id: number) {
    return `This action removes a #${id} account`;
  }
}
