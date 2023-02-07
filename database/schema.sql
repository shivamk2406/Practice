CREATE TABLE `test_practice` (
  `record_id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `id` varchar(36) NOT NULL,
  `name` varchar(1000) DEFAULT NULL,
  `subscription` varchar(1000) DEFAULT NULL,
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `creator` varchar(36) NOT NULL,
  `updater` varchar(36) NOT NULL,
  `active` tinyint(1) NOT NULL DEFAULT '1',
  PRIMARY KEY (`id`),
  UNIQUE KEY `record_id` (`record_id`)
) ENGINE=InnoDB AUTO_INCREMENT=101 DEFAULT CHARSET=utf8mb3