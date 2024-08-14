import { HttpException, HttpStatus, Injectable } from "@nestjs/common";

@Injectable()
export class TestService {
  getData(param): string {
    if (!param.uid) {
      throw new HttpException('必须包含id参数，并且id为数字', HttpStatus.BAD_REQUEST)
    }
    return 'get test data';
  }
}
