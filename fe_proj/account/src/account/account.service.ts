import { Injectable } from "@nestjs/common";
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

  login(loginAccountDto: LoginAccountDto) {
    return "This action logins a new account";
  }

  findAll() {
    return this.accountRepository.find();
  }

  findOne(id: number) {
    return this.accountRepository.findOneBy({ id });
  }

  findOneByUsername(username: string) {
    return this.accountRepository.findOneBy({ username });
  }

  update(id: number, updateAccountDto: UpdateAccountDto) {
    return `This action updates a #${id} account`;
  }

  remove(id: number) {
    return `This action removes a #${id} account`;
  }
}
