-- Modify "users" table
ALTER TABLE `users` DROP COLUMN `age`, ADD COLUMN `birthday` datetime(3) NULL;
