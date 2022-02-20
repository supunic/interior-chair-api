-- +migrate Up
CREATE TABLE IF NOT EXISTS `chairs`
(
    `id`         INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY COMMENT '椅子ID',
    `author_id`  INT UNSIGNED NOT NULL COMMENT '作者ID',
    `name`       VARCHAR(255) NOT NULL COMMENT '名前',
    `feature`    VARCHAR(255) NOT NULL COMMENT '特徴',
    `year`       INT          NOT NULL COMMENT '作成年',
    `image`      VARCHAR(255) NOT NULL COMMENT 'イメージ',
    `created_at` timestamp    NULL DEFAULT NULL,
    `updated_at` timestamp    NULL DEFAULT NULL
);

ALTER TABLE `chairs`
    ADD CONSTRAINT `chairs_author_id_foreign`
        FOREIGN KEY (`author_id`)
            REFERENCES `authors` (`id`)
            ON DELETE CASCADE
            ON UPDATE CASCADE;

-- +migrate Down
DROP TABLE IF EXISTS chairs;
