import { Column, Entity, PrimaryGeneratedColumn, Unique } from "typeorm";
@Entity('admin_user')
export class UserEntity{
  @PrimaryGeneratedColumn()
  id:number;

  @Column()
  @Unique(['username'])
  username:string;

  @Column()
  password:string;

  @Column()
  avatar:string;

  @Column()
  role:string;

  @Column()
  nickname:string;

  @Column()
  active:number;
}
