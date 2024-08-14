import { Body, Controller, Delete, Get, Param, ParseIntPipe, Post } from "@nestjs/common";
import { UserService } from "./user.service";
import { DeleteResult } from "typeorm";

@Controller('user')
export class UserController {
  constructor(private readonly userService: UserService) {
  }

  @Get(':id')
  getUser(@Param('id', ParseIntPipe) id) {
    console.log(id);
    console.log(typeof id);
    return this.userService.findOne(id)
  }

  @Get()
  getAllUser() {
    return this.userService.findAll()
  }

  @Post()
  create(@Body() body) {
    return this.userService.create(body)
  }

  @Delete(':id')
  remove(@Param('id', ParseIntPipe) id) {
    return this.userService.remove(id)
  }
}
