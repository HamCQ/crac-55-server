CREATE TABLE `crac_award` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `year` int NOT NULL DEFAULT '0',
  `callsign` varchar(45) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `continent` varchar(12) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `type` tinyint NOT NULL DEFAULT '0',
  `combination` int NOT NULL DEFAULT '0',
  `bncra_num` int NOT NULL DEFAULT '0',
  `img_url` varchar(2048) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `status` tinyint NOT NULL DEFAULT '0',
  `created_time` int NOT NULL DEFAULT '0',
  `update_time` int NOT NULL DEFAULT '0',
  `remark` varchar(45) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  PRIMARY KEY (`id`),
  KEY `crac_award_callsign_index` (`callsign`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci