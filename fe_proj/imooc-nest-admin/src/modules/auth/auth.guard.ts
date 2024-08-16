import { CanActivate, ExecutionContext, Injectable, UnauthorizedException } from "@nestjs/common";
import { Observable } from "rxjs";
import { Reflector } from "@nestjs/core";
import { IS_PUBLIC_KEY } from "./public.decorator";
import { JwtService } from "@nestjs/jwt";
import { JWT_SECRET_KEY } from "./auth.jwt.secret";

@Injectable()
export class AuthGuard implements CanActivate{
  constructor(
    private jwtService:JwtService,
    private reflector: Reflector) {}
  // 实现canActivate方法
  async canActivate(context: ExecutionContext): Promise<boolean> {
    // console.log('ExecutionContext', context);
    const isPublic = this.reflector.getAllAndOverride(IS_PUBLIC_KEY, [
      context.getHandler(),
      context.getClass(),// todo 这个是从哪里找呢
    ])
    if (isPublic) {
      return true;
    }

    const request = context.switchToHttp().getRequest();
    const token = extractTokenFromRequest(request)
    if (!token) {
      throw new UnauthorizedException();
    }
    try {
      const payload = await this.jwtService.verifyAsync(token, {
        secret: JWT_SECRET_KEY,
      })
      // console.log(payload);
      request['user'] = payload;
    }catch (e) {
      throw new UnauthorizedException();
    }
    return true;
  }

}


function extractTokenFromRequest(request) {
  const [type, token] = request.headers.authorization?.split(' ')??[];
  return type === "Bearer" ? token : '';
}
