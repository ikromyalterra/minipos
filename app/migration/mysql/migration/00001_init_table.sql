-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE `merchant` (
 `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
 `name` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL,
 `created_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
 `updated_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
 PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE `outlet` (
 `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
 `id_merchant` int(10) unsigned NOT NULL,
 `name` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL,
 `created_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
 `updated_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
 PRIMARY KEY (`id`),
 KEY `id_merchant` (`id_merchant`),
 CONSTRAINT `outlet_merchant` FOREIGN KEY (`id_merchant`) REFERENCES `merchant` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE `user` (
 `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
 `email` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL,
 `password` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
 `role` varchar(10) COLLATE utf8mb4_unicode_ci NOT NULL,
 `id_merchant` int(10) unsigned DEFAULT NULL,
 `id_outlet` int(10) unsigned DEFAULT NULL,
 `created_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
 `updated_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
 PRIMARY KEY (`id`),
 UNIQUE KEY `email` (`email`),
 KEY `id_merchant` (`id_merchant`),
 KEY `id_outlet` (`id_outlet`),
 CONSTRAINT `user_role_merchant` FOREIGN KEY (`id_merchant`) REFERENCES `merchant` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
 CONSTRAINT `user_role_outlet` FOREIGN KEY (`id_outlet`) REFERENCES `outlet` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE `user_token` (
 `id_token` int(10) unsigned NOT NULL,
 `id_user` int(10) unsigned NOT NULL,
 `created_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
 `updated_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
 PRIMARY KEY (`id_token`),
 KEY `id_user` (`id_user`),
 CONSTRAINT `user_has_user_token` FOREIGN KEY (`id_user`) REFERENCES `user` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE `product` (
 `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
 `id_merchant` int(10) unsigned NOT NULL,
 `sku` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL,
 `image` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
 `name` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL,
 `price` int(10) unsigned NOT NULL DEFAULT 0,
 `created_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
 `updated_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
 PRIMARY KEY (`id`),
 KEY `id_merchant` (`id_merchant`),
 KEY `sku` (`sku`),
 CONSTRAINT `product_merchant` FOREIGN KEY (`id_merchant`) REFERENCES `merchant` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE `product_outlet` (
 `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
 `id_product` int(10) unsigned NOT NULL,
 `id_outlet` int(10) unsigned NOT NULL,
 `price` int(10) unsigned NOT NULL DEFAULT 0,
 `created_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
 `updated_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
 PRIMARY KEY (`id`),
 UNIQUE KEY `product_outlet` (`id_product`,`id_outlet`) USING BTREE,
 KEY `id_outlet` (`id_outlet`),
 CONSTRAINT `product_outlet_product` FOREIGN KEY (`id_product`) REFERENCES `product` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
 CONSTRAINT `outlet_has_product_outlet` FOREIGN KEY (`id_outlet`) REFERENCES `outlet` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE `product_outlet`;
DROP TABLE `product`;
DROP TABLE `user_token`;
DROP TABLE `user`;
DROP TABLE `outlet`;
DROP TABLE `merchant`;
