CREATE DATABASE IF NOT EXISTS `go_db`;

USE `go_db`;

CREATE TABLE IF NOT EXISTS characters(
    `id` INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
    `name` VARCHAR(256) NOT NULL,
    `rality` INT NOT NULL,
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP `updated_at` TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE = INNODB DEFAULT CHARSET = utf8mb4;

INSERT INTO
    characters(name, reality) VALUE('佐藤文哉', 5);

INSERT INTO
    characters(name, reality) VALUE('菅さん', 5);

INSERT INTO
    characters(name, reality) VALUE('野田さん', 5);

INSERT INTO
    characters(name, reality) VALUE('谷さん', 5);

INSERT INTO
    characters(name, reality) VALUE('satofumi', 4);

INSERT INTO
    characters(name, reality) VALUE('雪だるま', 4);

INSERT INTO
    characters(name, reality) VALUE('竈門炭治郎', 4);

INSERT INTO
    characters(name, reality) VALUE('チェンソーマン', 4);

INSERT INTO
    characters(name, reality) VALUE('ドラえもん', 4);

INSERT INTO
    characters(name, reality) VALUE('善逸', 3);

INSERT INTO
    characters(name, reality) VALUE('のび太くん', 3);

INSERT INTO
    characters(name, reality) VALUE('キリン', 3);