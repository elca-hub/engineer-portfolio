-- Modify "users" table
ALTER TABLE `users` MODIFY COLUMN `id` bigint unsigned NOT NULL AUTO_INCREMENT, ADD COLUMN `icon_path` longtext NULL, ADD COLUMN `header_path` longtext NULL, ADD COLUMN `bio_path` longtext NULL;
-- Create "external_service_urls" table
CREATE TABLE `external_service_urls` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL,
  `updated_at` datetime(3) NULL,
  `deleted_at` datetime(3) NULL,
  `name` varchar(100) NOT NULL,
  `url` varchar(255) NOT NULL,
  `user_id` bigint unsigned NULL,
  PRIMARY KEY (`id`),
  INDEX `fk_users_external_service_urls` (`user_id`),
  INDEX `idx_external_service_urls_deleted_at` (`deleted_at`),
  CONSTRAINT `fk_users_external_service_urls` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON UPDATE CASCADE ON DELETE CASCADE
) CHARSET utf8mb4 COLLATE utf8mb4_0900_ai_ci;
-- Create "skills" table
CREATE TABLE `skills` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL,
  `updated_at` datetime(3) NULL,
  `deleted_at` datetime(3) NULL,
  `name` varchar(255) NOT NULL,
  `status` varchar(255) NOT NULL,
  `when` datetime(3) NOT NULL,
  `sort_index` bigint NOT NULL,
  `user_id` bigint unsigned NULL,
  PRIMARY KEY (`id`),
  INDEX `fk_users_skills` (`user_id`),
  INDEX `idx_skills_deleted_at` (`deleted_at`),
  CONSTRAINT `fk_users_skills` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON UPDATE CASCADE ON DELETE CASCADE
) CHARSET utf8mb4 COLLATE utf8mb4_0900_ai_ci;
