import { Body, Controller, Delete, Get, Param, ParseIntPipe, Post, Req } from "@nestjs/common";
import { UserService } from "./user.service";
import { DeleteResult } from "typeorm";
import { wrapperResponse } from "../../utils";
import { request } from "express";

@Controller('user')
export class UserController {
  constructor(private readonly userService: UserService) {
  }

  @Get('info')
  getUserByToken(@Req() request) {
    return wrapperResponse(
      this.userService.findByUsername(request.user.username),
      '获取用户信息成功',
    )
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
