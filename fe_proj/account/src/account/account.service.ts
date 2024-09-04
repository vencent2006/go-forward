import { HttpException, HttpStatus, Injectable } from "@nestjs/common";
import { CreateAccountDto } from "./dto/create-account.dto";
import { UpdateAccountDto } from "./dto/update-account.dto";
import { LoginAccountByMailDto, LoginAccountDto } from "./dto/login-account.dto";
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
    const account = await this.findOneByUsername(loginAccountDto.username);
    // todo use md5 encrypted password
    if (!account || account.password != loginAccountDto.password) {
      throw new HttpException("invalid login account", HttpStatus.BAD_REQUEST);
    }
    return account;
  }

  async loginByMail(loginAccountDto: LoginAccountByMailDto) {
    if (loginAccountDto.mail == "") {
      throw new HttpException("mail is null", HttpStatus.BAD_REQUEST);
    }
    const account = await this.findOneByMail(loginAccountDto.mail);
    // todo use md5 encrypted password
    if (!account || account.password != loginAccountDto.password) {
      throw new HttpException("invalid login account", HttpStatus.BAD_REQUEST);
    }
    return account;
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

  findOneByMail(mail: string): Promise<AccountEntity> {
    return this.accountRepository.findOneBy({ mail });
  }

  update(id: number, updateAccountDto: UpdateAccountDto) {
    return `This action updates a #${id} account`;
  }

  remove(id: number) {
    return `This action removes a #${id} account`;
  }
}
