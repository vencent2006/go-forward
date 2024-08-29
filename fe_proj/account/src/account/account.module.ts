import { Module } from "@nestjs/common";
import { AccountService } from "./account.service";
import { AccountController } from "./account.controller";
import { TypeOrmModule } from "@nestjs/typeorm";
import { AccountEntity } from "./entities/account.entity";

@Module({
  imports: [TypeOrmModule.forFeature([AccountEntity])],
  controllers: [AccountController],
  providers: [AccountService]
})
export class AccountModule {
}
