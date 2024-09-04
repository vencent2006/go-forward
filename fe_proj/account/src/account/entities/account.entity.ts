import { Column, CreateDateColumn, Entity, PrimaryGeneratedColumn, Unique, UpdateDateColumn } from "typeorm";

@Entity("account")
export class AccountEntity {
  @PrimaryGeneratedColumn()
  id: number;

  @Column()
  @Unique(["username"])
  username: string;

  @Column()
  password: string;

  @Column()
  avatar: string;

  @Column()
  role: number;

  @Column()
  nickname: string;

  @Column()
  mail: string;

  @Column()
  status: number;

  @CreateDateColumn() // 创建时自动插入
  create_time: Date;

  @UpdateDateColumn() // 更新时自动更新
  update_time: Date;
}

/**
 * 账号状态
 */
export enum AccountStatus {
  VALID = 1,
  INVALID = 2,
}

/**
 * 账号角色
 */
export enum AccountRole {
  USER = 1,
  ADMIN = 2,
}
