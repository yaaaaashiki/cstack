-- +migrate Up
TRUNCATE TABLE `gorp_migrations`;

DROP TABLE IF EXISTS `items`;

CREATE TABLE `items` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` varchar(255) DEFAULT '',
  `name` varchar(255) NOT NULL DEFAULT '',
  `price` bigint(255)  NOT NULL DEFAULT '',
  `current_payment_price` bigint(255) NOT NULL DEFAULT '',
  `icon_image` varchar(255) DEFAULT '',
  `description` varchar(255) DEFAULT '',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  CONSTRAINT `items_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
