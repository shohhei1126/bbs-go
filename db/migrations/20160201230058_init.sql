-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE `users` (
  `id`           INT(10) UNSIGNED                        NOT NULL AUTO_INCREMENT,
  `username`     VARCHAR(64) COLLATE utf8mb4_unicode_ci  NOT NULL,
  `password`     VARCHAR(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `display_name` VARCHAR(188) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `status`       TINYINT(3) UNSIGNED                     NOT NULL DEFAULT '1',
  `created_at`   DATETIME                                NOT NULL,
  `updated_at`   DATETIME                                NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `username` (`username`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_unicode_ci;

CREATE TABLE `threads` (
  `id`            INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `user_id`       INT(10) UNSIGNED NOT NULL,
  `title`         TEXT             NOT NULL,
  `body`          TEXT             NOT NULL,
  `created_at`    DATETIME         NOT NULL,
  `updated_at`    DATETIME         NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_unicode_ci;

INSERT INTO users (id, username, password, display_name, status, created_at, updated_at) VALUES (1, 'hoge', 'hoge', 'hoge', 1, '2016-02-16 21:13:57', '2016-02-16 21:13:58');
INSERT INTO users (id, username, password, display_name, status, created_at, updated_at) VALUES (2, 'fuga', 'fuga', 'fuga', 1, '2016-02-16 21:14:34', '2016-02-16 21:14:35');
INSERT INTO users (id, username, password, display_name, status, created_at, updated_at) VALUES (3, 'bar', 'bar', 'bar', 1, '2016-02-16 21:14:50', '2016-02-16 21:14:52');
INSERT INTO users (id, username, password, display_name, status, created_at, updated_at) VALUES (4, 'buzz', 'buzz', 'buzz', 1, '2016-02-16 21:15:22', '2016-02-16 21:15:24');
INSERT INTO threads (id, user_id, title, body, created_at, updated_at) VALUES (1, 1, 'a', 'a', '2016-02-16 21:19:36', '2016-02-16 21:19:38');
INSERT INTO threads (id, user_id, title, body, created_at, updated_at) VALUES (2, 2, 'b', 'b', '2016-02-16 21:19:51', '2016-02-18 21:19:52');
INSERT INTO threads (id, user_id, title, body, created_at, updated_at) VALUES (3, 3, 'c', 'c', '2016-02-16 21:20:01', '2016-02-20 21:20:03');
INSERT INTO threads (id, user_id, title, body, created_at, updated_at) VALUES (4, 4, 'd', 'd', '2016-02-16 21:20:12', '2016-02-16 21:20:13');
INSERT INTO threads (id, user_id, title, body, created_at, updated_at) VALUES (5, 1, 'e', 'e', '2016-02-17 10:36:01', '2016-02-17 10:36:03');
INSERT INTO threads (id, user_id, title, body, created_at, updated_at) VALUES (6, 2, 'f', 'f', '2016-02-17 10:36:19', '2016-02-15 10:36:21');
INSERT INTO threads (id, user_id, title, body, created_at, updated_at) VALUES (7, 3, 'g', 'g', '2016-02-17 10:36:30', '2016-04-17 10:36:32');
INSERT INTO threads (id, user_id, title, body, created_at, updated_at) VALUES (8, 4, 'h', 'h', '2016-02-17 10:36:41', '2016-01-17 10:36:43');
INSERT INTO threads (id, user_id, title, body, created_at, updated_at) VALUES (9, 1, 'i', 'i', '2016-02-17 10:36:53', '2016-04-17 10:36:55');
INSERT INTO threads (id, user_id, title, body, created_at, updated_at) VALUES (10, 2, 'j', 'j', '2016-02-17 10:37:16', '2016-02-17 10:37:18');
INSERT INTO threads (id, user_id, title, body, created_at, updated_at) VALUES (11, 3, 'k', 'k', '2016-02-17 10:37:30', '2016-02-17 10:37:32');
INSERT INTO threads (id, user_id, title, body, created_at, updated_at) VALUES (12, 4, 'l', 'l', '2016-02-17 10:37:44', '2016-02-21 10:37:45');
INSERT INTO threads (id, user_id, title, body, created_at, updated_at) VALUES (13, 1, 'm', 'm', '2016-02-17 10:37:56', '2016-02-17 10:37:59');

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE `users`;
DROP TABLE `threads`;