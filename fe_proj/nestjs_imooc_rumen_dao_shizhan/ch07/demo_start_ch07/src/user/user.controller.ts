import {
  Controller,
  Get,
  Post,
  Body,
  Patch,
  Param,
  Delete,
} from '@nestjs/common';
import { UserService } from './user.service';
import { CreateUserDto } from './dto/create-user.dto';
import { UpdateUserDto } from './dto/update-user.dto';
import { ConfigService } from '@nestjs/config';
import { ConfigEnum } from 'src/enum/const';

@Controller('user')
export class UserController {
  constructor(
    private readonly userService: UserService,
    private configService: ConfigService,
  ) {}

  @Post()
  create(@Body() createUserDto: CreateUserDto) {
    return this.userService.create(createUserDto);
  }

  @Get()
  findAll() {
    // const db = this.configService.get(ConfigEnum.DB);
    // const db = this.configService.get('db');
    // console.log(db);
    const dbPort = this.configService.get(ConfigEnum.DB_PORT);
    console.log(dbPort);
    return this.userService.findAll();
  }

  @Get('/profile/:id')
  getUserProfile(@Param('id') id: number): any {
    console.log('controller getUserProfile | id = ', id);
    return this.userService.findProfile(id);
  }

  @Get('/logs/:id')
  getUserLogs(@Param('id') id: number): any {
    console.log('controller getUserLogs | id = ', id);
    return this.userService.findUserLogs(id);
  }

  @Get(':id')
  findOne(@Param('id') id: string) {
    return this.userService.findOne(+id);
  }

  @Patch(':id')
  update(@Param('id') id: string, @Body() updateUserDto: UpdateUserDto) {
    return this.userService.update(+id, updateUserDto);
  }

  @Delete(':id')
  remove(@Param('id') id: string) {
    return this.userService.remove(+id);
  }


}
