
-- +migrate Up
CREATE TABLE IF NOT EXISTS `games` (
    `id` INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY COMMENT 'ゲームID',
    `game_type_id` INT UNSIGNED NOT NULL COMMENT 'ゲームタイプID',
    `score` INT NOT NULL COMMENT '得点'
);

ALTER TABLE `games`
    ADD CONSTRAINT `games_game_type_id_foreign`
        FOREIGN KEY (`game_type_id`)
        REFERENCES `game_types` (`id`)
        ON DELETE CASCADE
        ON UPDATE CASCADE;

-- +migrate Down
DROP TABLE IF EXISTS games;
