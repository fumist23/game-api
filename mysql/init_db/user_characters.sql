CREATE DATABASE IF NOT EXISTS `go_db`;

USE `go_db`;

CREATE TABLE IF NOT EXISTS user_characters(
    `id` INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
    `userId` INT NOT NULL,
    `characterId` INT NOT NULL,
    `characterName` VARCHAR(256) NOT NULL
) ENGINE = INNODB DEFAULT CHARSET = utf8mb4;