import { ArgumentsHost, Catch, ExceptionFilter, HttpException } from "@nestjs/common";
import { Request, Response } from "express";

@Catch(HttpException)
export class HttpFilter implements ExceptionFilter {
  catch(exception: HttpException, host: ArgumentsHost) {
    const ctx = host.switchToHttp()
    const request = ctx.getRequest<Request>()
    const response = ctx.getResponse<Response>()
    const status = exception.getStatus()
    response.status(status).json({
      success: false,
      time: new Date(),
      message: exception.getResponse()['message']? exception.getResponse()['message']: exception.message,
      status,
      path: request.url
    })
  }
}
