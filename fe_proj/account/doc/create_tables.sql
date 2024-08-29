CREATE TABLE `account` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `username` varchar(255) NOT NULL COMMENT '用户名',
  `password` varchar(255) NOT NULL COMMENT '加密后的密码',
  `avatar` varchar(255) NOT NULL COMMENT '头像',
  `role` int(10) NOT NULL COMMENT '角色：0，普通用户；1，管理员',
  `nickname` varchar(255) NOT NULL COMMENT '昵称',
  `status` tinyint NOT NULL COMMENT '状态：0，失效；1，生效',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `UNI_USERNAME` (`username`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
