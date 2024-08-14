import { Controller, Get, Param, ParseIntPipe } from "@nestjs/common";

@Controller('user')
export class UserController {
  @Get(':id')
  getUser(@Param('id', ParseIntPipe) id): string {
    console.log(id);
    console.log(typeof id);
    console.log('hello');
    return 'get user data: ' + id;
  }
}
