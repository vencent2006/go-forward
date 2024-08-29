import { Injectable } from "@nestjs/common";
import { CreateAccountDto } from "./dto/create-account.dto";
import { UpdateAccountDto } from "./dto/update-account.dto";
import { LoginAccountDto } from "./dto/login-account.dto";
import { InjectRepository } from "@nestjs/typeorm";
import { DeleteResult, Repository } from "typeorm";
import { AccountEntity } from "./entities/account.entity";

@Injectable()
export class AccountService {
  constructor(
    @InjectRepository(AccountEntity) // 注入repository
    private readonly accountRepository: Repository<AccountEntity>
  ) {
  }

  create(createAccountDto: CreateAccountDto) {
    // 1. username 是否存在
    return "This action adds a new account";
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

  update(id: number, updateAccountDto: UpdateAccountDto) {
    return `This action updates a #${id} account`;
  }

  remove(id: number) {
    return `This action removes a #${id} account`;
  }
}
