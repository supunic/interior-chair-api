
-- +migrate Up
CREATE TABLE IF NOT EXISTS `game_types` (
    `id` INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY COMMENT 'ゲームタイプID',
    `name` VARCHAR(255) NOT NULL COMMENT '名前'
);

-- +migrate Down
DROP TABLE IF EXISTS game_types;
