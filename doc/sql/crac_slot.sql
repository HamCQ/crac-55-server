CREATE TABLE `crac_slot` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `user_id` int unsigned NOT NULL,
  `bncra` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `callsign` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `year` int NOT NULL DEFAULT '0',
  `mode` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `band` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `status` tinyint NOT NULL DEFAULT '0',
  `created_time` int NOT NULL DEFAULT '0',
  `update_time` int NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  KEY `crac_slot_user_id_index` (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci