-- +migrate Up
CREATE TABLE IF NOT EXISTS `chairs`
(
    `id`         INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY COMMENT '椅子ID',
    `author_id`  INT UNSIGNED NOT NULL COMMENT '椅子作者ID',
    `name`       VARCHAR(255) NOT NULL COMMENT '名前',
    `feature`    VARCHAR(255) NOT NULL COMMENT '特徴',
    `year`       INT          NOT NULL COMMENT '作成年',
    `image`      VARCHAR(255) NOT NULL COMMENT 'イメージ',
    `created_at` timestamp    NULL DEFAULT NULL,
    `updated_at` timestamp    NULL DEFAULT NULL
);

-- +migrate Down
DROP TABLE IF EXISTS chairs;
