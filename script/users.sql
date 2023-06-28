
DROP TABLE IF EXISTS `users`;

CREATE TABLE `livestreamgo`.`users` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `username` VARCHAR(225) NULL,
  `password` VARCHAR(225) NULL,
  `created_at` DATETIME NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` DATETIME NULL,
  PRIMARY KEY (`id`));