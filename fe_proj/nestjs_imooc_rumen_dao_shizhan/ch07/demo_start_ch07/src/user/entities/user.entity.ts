import { Logs } from "src/logs/logs.entity";
import { Roles } from "src/roles/roles.entity";
import { Column, Entity, OneToMany, PrimaryGeneratedColumn, ManyToMany, JoinTable, OneToOne } from "typeorm";
import { Profile } from "./profile.entity";
import { profile } from "console";

@Entity()
export class User {
  @PrimaryGeneratedColumn()
  id: number;

  @Column()
  username: string;

  @Column()
  password: string;

  // typescript 与 数据库的关联关系 mapping
  @OneToMany(()=>Logs, (logs)=>logs.user)
  logs: Logs[];

  @ManyToMany(()=>Roles, (roles)=>roles.users)
  @JoinTable({name: 'users_roles'})
  roles: Roles[];

  @OneToOne(()=> Profile, (profile)=>profile.user)
  profile: Profile;
}
