import { Controller, Get, Param, ParseIntPipe } from "@nestjs/common";

@Controller('book')
export class BookController {
  @Get(':id')
  getBookById(@Param('id', ParseIntPipe) id) {
    return 'book id: '+id;
  }
}
