import {Injectable} from "@nestjs/common";
import { InjectRepository } from "@nestjs/typeorm";
import { UserEntity } from "./user.entity";
import { DeleteResult, Repository } from "typeorm";
import { CreateUserDto } from "./create-user.dto";
@Injectable()
export class UserService{
  constructor(
    @InjectRepository(UserEntity) // 注入repository
    private readonly userRepository:Repository<UserEntity>
  ) {}

  findOne(id:number):Promise<UserEntity>{
    return this.userRepository.findOneBy({id});
  }

  findAll():Promise<UserEntity[]>{
    return this.userRepository.find();
  }

  create(createUserDto:CreateUserDto):Promise<UserEntity>{
    const user = new UserEntity();
    user.username = createUserDto.username;
    user.password = createUserDto.password;
    user.role = createUserDto.role;
    user.avatar = createUserDto.avatar;
    user.nickname = createUserDto.nickname;
    user.active = 1;
    return this.userRepository.save(user);
  }

  remove(id:number):Promise<DeleteResult>{
    return this.userRepository.delete(id);
  }

  findByUsername(username:string):Promise<UserEntity|null>{
    return this.userRepository.findOneBy({username})
  }
}
