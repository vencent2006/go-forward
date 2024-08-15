import { Module } from '@nestjs/common';
import { UserController } from './user.controller';
import { TypeOrmModule } from "@nestjs/typeorm";
import { UserEntity } from "./user.entity";
import { UserService } from "./user.service";

@Module({
  imports: [TypeOrmModule.forFeature([UserEntity])],
  controllers: [UserController],
  providers: [UserService],
  exports: [UserService],// 说明 UserService可以被别的module引用
})
export class UserModule {}
