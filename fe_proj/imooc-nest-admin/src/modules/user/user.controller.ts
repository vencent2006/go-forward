import { Controller, Get, Param, ParseIntPipe } from "@nestjs/common";
import { UserService } from "./user.service";

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
}
