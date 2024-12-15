import { Injectable } from '@nestjs/common';
import { CreateUserDto } from './dto/create-user.dto';
import { UpdateUserDto } from './dto/update-user.dto';
import { InjectRepository } from '@nestjs/typeorm';
import { User } from './entities/user.entity';
import { Repository } from 'typeorm';
import { log } from 'console';
import { Logs } from 'src/logs/logs.entity';

@Injectable()
export class UserService {
  constructor(
    @InjectRepository(User) private readonly userRepository: Repository<User>,
    @InjectRepository(Logs) private readonly logsRepository: Repository<Logs>,
  ) { }
  
  create(createUserDto: CreateUserDto) {
    return 'This action adds a new user';
  }

  find(username: string) {
    return this.userRepository.findOne({ where: { username } });
  }

  findAll() {
    console.log('*********')
    return this.userRepository.find();
  }

  findOne(id: number) {
    return this.userRepository.findOne({ where: { id } });
  }

  findProfile(id: number) {
    console.log('id = ', id)
    return this.userRepository.findOne({
      where:{
        id,
      },
      relations: {
        profile: true,
      },
    })
  }

  async findUserLogs(id: number) {
    const user = await this.findOne(id)
    return this.logsRepository.findOne({
      where:{
        user,
      },
      relations: {
        user: true,
      },
    })
  }

  update(id: number, updateUserDto: UpdateUserDto) {
    return `This action updates a #${id} user`;
  }

  remove(id: number) {
    return `This action removes a #${id} user`;
  }
}
