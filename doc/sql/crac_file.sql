CREATE TABLE `crac_file` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `user_id` int unsigned NOT NULL,
  `status` tinyint NOT NULL DEFAULT '0',
  `sys_filename` varchar(512) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `filename` varchar(512) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `qso_num` int NOT NULL DEFAULT '0',
  `remark` varchar(128) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `created_time` int NOT NULL DEFAULT '0',
  `update_time` int NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  KEY `crac_file_user_id_index` (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci