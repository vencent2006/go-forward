import { Module } from "@nestjs/common";
import { AppController } from "./app.controller";
import { AppService } from "./app.service";
import { AccountModule } from "./account/account.module";
import { TypeOrmModule } from "@nestjs/typeorm";
import { AuthModule } from "./auth/auth.module";
import { ConfigModule } from "@nestjs/config";
import { connectionOptions } from "./config/orm.config";


@Module({
  imports: [
    ConfigModule.forRoot({
      isGlobal: true
    }),
    TypeOrmModule.forRoot(connectionOptions),
    AccountModule,
    AuthModule],
  controllers: [AppController],
  providers: [AppService]
})
export class AppModule {
}
