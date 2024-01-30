CREATE TABLE `users`
(
    `id`          bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'PK',
    `name`        varchar(128) NOT NULL DEFAULT '' COMMENT 'User name',
    `gender`      int(8) NOT NULL DEFAULT 0 COMMENT 'User gender',
    `age`         int(64) NOT NULL DEFAULT 0 COMMENT 'User age',
    `introduce`   Text NULL COMMENT 'User introduce',
    `created_at` timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'User information create time',
    `updated_at` timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'User information update time',
    `deleted_at` timestamp NULL DEFAULT NULL COMMENT 'User information delete time',
    PRIMARY KEY (`id`),
    KEY           `idx_name` (`name`,`deleted_at`) COMMENT 'User name index'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='User information table'