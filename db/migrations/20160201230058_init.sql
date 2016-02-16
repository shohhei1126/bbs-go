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



-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE `users`;
DROP TABLE `threads`;