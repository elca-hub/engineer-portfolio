-- Modify "users" table
ALTER TABLE `users` MODIFY COLUMN `id` varchar(255) NOT NULL, ADD UNIQUE INDEX `uni_users_id` (`id`);
-- Modify "external_service_urls" table
ALTER TABLE `external_service_urls` MODIFY COLUMN `user_id` varchar(255) NULL;
-- Modify "skills" table
ALTER TABLE `skills` MODIFY COLUMN `user_id` varchar(255) NULL;