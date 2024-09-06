import { TypeOrmModuleOptions } from "@nestjs/typeorm";
import * as fs from "fs";
import * as dotenv from "dotenv";
import { DB_Enum } from "src/enums/typeorm-config.enum";
import * as path from "path";


// 匹配所有实体
const entitiesDir = [path.join(__dirname, "../**/*.entity{.ts,.js}")];

// 通过环境变量读取不同的.env文件
function getEnv(env: string): Record<string, string> {
  if (fs.existsSync(env)) {
    return dotenv.parse(fs.readFileSync(env));
  }
  return {};
}

// 通过dotENV来解析不同的配置
function buildConnectionOptions() {
  const defaultConfig = getEnv(".env");
  const envConfig = getEnv(`.env.${process.env.NODE_ENV || "development"}`);
  const config = {
    ...defaultConfig,
    ...envConfig
  };
  return {
    type: "mysql",
    host: config[DB_Enum.HOST],
    port: (config[DB_Enum.PORT]),
    username: config[DB_Enum.USERNAME],
    password: config[DB_Enum.PASSWORD],
    database: config[DB_Enum.DATABASE],
    entities: entitiesDir,
    synchronize: JSON.parse(config[DB_Enum.SYNC]),//同步本地的schema与数据库 ->初始化的时候使用
    // logging: ["error"],//日志等级
    autoLoadEntities: true
  };
}

export const connectionOptions = buildConnectionOptions() as TypeOrmModuleOptions;
