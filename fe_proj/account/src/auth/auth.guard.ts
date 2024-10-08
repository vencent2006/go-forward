import { CanActivate, ExecutionContext, Injectable, UnauthorizedException } from "@nestjs/common";
import { Observable } from "rxjs";
import { JwtService } from "@nestjs/jwt";
import { Reflector } from "@nestjs/core";
import { IS_PUBLIC_KEY } from "./public.decorator";
import { JWT_SECRET_KEY } from "./auth.jwt.secret";

@Injectable()
export class AuthGuard implements CanActivate {
  constructor(
    private jwtService: JwtService,
    private reflector: Reflector) {
  }

  async canActivate(
    context: ExecutionContext
  ): Promise<boolean> {
    // console.log('ExecutionContext', context);
    const isPublic = this.reflector.getAllAndOverride(IS_PUBLIC_KEY, [
      context.getHandler(),
      context.getClass()// todo 这个是从哪里找呢
    ]);
    if (isPublic) {
      // 公开的要放行
      return true;
    }

    const request = context.switchToHttp().getRequest();
    const token = extractTokenFromRequest(request);
    if (!token) {
      throw new UnauthorizedException();
    }
    try {
      const payload = await this.jwtService.verifyAsync(token, {
        secret: JWT_SECRET_KEY
      });
      // console.log(payload);
      request["user"] = payload;
    } catch (e) {
      throw new UnauthorizedException("token verification failed");
    }
    return true;
  }

}


function extractTokenFromRequest(request) {
  const [type, token] = request.headers.authorization?.split(" ") ?? [];
  return type === "Bearer" ? token : "";
}
