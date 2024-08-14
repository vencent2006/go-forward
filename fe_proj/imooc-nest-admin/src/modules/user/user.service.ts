import {Injectable} from "@nestjs/common";
import { InjectRepository } from "@nestjs/typeorm";
import { UserEntity } from "./user.entity";
import { Repository } from "typeorm";
@Injectable()
export class UserService{
  constructor(
    @InjectRepository(UserEntity)
    private readonly userRepository:Repository<UserEntity>
  ) {}

  findOne(id:number):Promise<UserEntity>{
    return this.userRepository.findOneBy({id});
  }

}
