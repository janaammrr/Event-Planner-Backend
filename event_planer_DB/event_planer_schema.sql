CREATE DATABASE IF NOT EXISTS `EventPlanner`
DEFAULT CHARACTER SET utf8mb4
COLLATE utf8mb4_unicode_ci;

USE `EventPlanner`;

CREATE TABLE IF NOT EXISTS `users` (
`user_id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
`name` VARCHAR(100) NOT NULL,
`email` VARCHAR(255) NOT NULL,
`password_hash` VARCHAR(255) NOT NULL,
`role` ENUM('organizer','attendee') NOT NULL DEFAULT 'attendee',
`created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
PRIMARY KEY (`user_id`),
UNIQUE KEY `ux_users_email` (`email`),
KEY `ix_users_created_at` (`created_at`)
) ENGINE=InnoDB
DEFAULT CHARSET=utf8mb4
COLLATE=utf8mb4_unicode_ci;

