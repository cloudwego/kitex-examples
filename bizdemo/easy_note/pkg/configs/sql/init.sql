CREATE TABLE `user`
(
    `id`         bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'PK',
    `user_name`  varchar(128) NOT NULL DEFAULT '' COMMENT 'UserName',
    `password`   varchar(128) NOT NULL DEFAULT '' COMMENT 'Password',
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'User account create time',
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'User account update time',
    `deleted_at` timestamp NULL DEFAULT NULL COMMENT 'User account delete time',
    PRIMARY KEY (`id`),
    KEY          `idx_user_name` (`user_name`, `deleted_at`) COMMENT 'UserName index'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='User account table';

CREATE TABLE `note`
(
    `id`         bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'PK',
    `user_id`    int(64) NOT NULL DEFAULT 0 COMMENT 'UserID',
    `title`      varchar(128) NOT NULL DEFAULT '' COMMENT 'Title',
    `content`    TEXT NULL COMMENT 'Content',
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Note creates time',
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'Note updates time',
    `deleted_at` timestamp NULL DEFAULT NULL COMMENT 'Note deletes time',
    PRIMARY KEY (`id`),
    KEY          `idx_user_id` (`user_id`, `deleted_at`) COMMENT 'UserID index'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='Note table';