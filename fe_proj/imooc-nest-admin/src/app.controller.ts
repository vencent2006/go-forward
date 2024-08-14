import { Body, Controller, Get, Param, Post, UseFilters } from "@nestjs/common";
import { AppService } from "./app.service";
import { TestService } from "./test.service";
import { HttpExceptionFilter } from "./exception/http-exception.filter";

@Controller()
export class AppController {
  constructor(
    private readonly appService: AppService,
    private readonly testService: TestService,
  ) {}

  @Get()
  getHello(): string {
    return this.appService.getHello();
  }

  @Get('/data/:id')
  @UseFilters(new HttpExceptionFilter())
  getData(@Param() params): string {
    console.log(params);
    return this.testService.getData(params);
  }

  @Post('data')
  saveData(@Body() data): string {
    console.log(data);
    return 'add data ' + JSON.stringify(data);
  }
}
