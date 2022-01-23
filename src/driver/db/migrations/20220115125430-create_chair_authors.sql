-- +migrate Up
CREATE TABLE IF NOT EXISTS `chair_authors`
(
    `id`          INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY COMMENT '椅子作者ID',
    `name`        VARCHAR(255) NOT NULL COMMENT '名前',
    `description` VARCHAR(255) NOT NULL COMMENT '説明',
    `birth_year`  INT          NOT NULL COMMENT '誕生年',
    `died_year`   INT COMMENT '没年',
    `image`       VARCHAR(255) NOT NULL COMMENT 'イメージ',
    `created_at`  timestamp    NULL DEFAULT NULL,
    `updated_at`  timestamp    NULL DEFAULT NULL
);

-- +migrate Down
DROP TABLE IF EXISTS chair_authoers;
