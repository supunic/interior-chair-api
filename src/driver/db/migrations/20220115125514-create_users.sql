
-- +migrate Up
CREATE TABLE IF NOT EXISTS `users` (
    `id` INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY COMMENT 'ユーザーID',
    `name` VARCHAR(255) NOT NULL COMMENT '名前'
);

ALTER TABLE `users`
    ADD UNIQUE KEY `users_name_unique` (`name`);

-- +migrate Down
DROP TABLE IF EXISTS users;
