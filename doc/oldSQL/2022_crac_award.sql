CREATE TABLE `award` (
  `id` int NOT NULL AUTO_INCREMENT,
  `call_sign` varchar(45) NOT NULL DEFAULT '',
  `continent` varchar(12) DEFAULT '' COMMENT '所属板块',
  `type` tinyint NOT NULL DEFAULT '0' COMMENT '0位置 1金奖 2银奖 3铜奖',
  `combination` int DEFAULT '0' COMMENT '组合数',
  `bncra_num` int DEFAULT '0' COMMENT 'BnCRA数',
  `year` int NOT NULL DEFAULT '0' COMMENT '隶属哪一年的活动',
  `img_url` varchar(2048) NOT NULL DEFAULT '' COMMENT '托管在oss的图片地址',
  `created_time` int NOT NULL DEFAULT '0',
  `update_time` int DEFAULT '0',
  `remark` varchar(45) DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='奖状'