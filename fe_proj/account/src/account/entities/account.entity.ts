import { Entity, PrimaryGeneratedColumn, Column, Unique, CreateDateColumn, UpdateDateColumn } from "typeorm";

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
  status: number;

  @CreateDateColumn() // 创建时自动插入
  create_time: Date;

  @UpdateDateColumn() // 更新时自动更新
  update_time: Date;
}
