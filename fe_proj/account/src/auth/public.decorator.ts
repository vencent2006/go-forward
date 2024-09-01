import { SetMetadata } from "@nestjs/common";

export const IS_PUBLIC_KEY = 'IS_PUBLIC_KEY'
// public注解 标志为公共路径
export const Public = () => SetMetadata(IS_PUBLIC_KEY, true)

