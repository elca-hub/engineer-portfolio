-- Step 1: Check and drop foreign key constraints if they exist

-- Check for external_service_urls foreign key
SET @fk_exists_ext = (SELECT COUNT(*) FROM INFORMATION_SCHEMA.TABLE_CONSTRAINTS 
                      WHERE TABLE_SCHEMA = DATABASE() 
                      AND TABLE_NAME = 'external_service_urls' 
                      AND CONSTRAINT_TYPE = 'FOREIGN KEY' 
                      AND CONSTRAINT_NAME = 'fk_users_external_service_urls');

SET @drop_fk_ext = CONCAT('ALTER TABLE `external_service_urls` DROP FOREIGN KEY `fk_users_external_service_urls`');

-- Only drop if it exists
SET @sql_ext = IF(@fk_exists_ext > 0, @drop_fk_ext, 'SELECT 1');
PREPARE stmt FROM @sql_ext;
EXECUTE stmt;
DEALLOCATE PREPARE stmt;

-- Check for skills foreign key
SET @fk_exists_skills = (SELECT COUNT(*) FROM INFORMATION_SCHEMA.TABLE_CONSTRAINTS 
                         WHERE TABLE_SCHEMA = DATABASE() 
                         AND TABLE_NAME = 'skills' 
                         AND CONSTRAINT_TYPE = 'FOREIGN KEY' 
                         AND CONSTRAINT_NAME = 'fk_users_skills');

SET @drop_fk_skills = CONCAT('ALTER TABLE `skills` DROP FOREIGN KEY `fk_users_skills`');

-- Only drop if it exists
SET @sql_skills = IF(@fk_exists_skills > 0, @drop_fk_skills, 'SELECT 1');
PREPARE stmt FROM @sql_skills;
EXECUTE stmt;
DEALLOCATE PREPARE stmt;

-- Step 2: Modify column types in the parent table
ALTER TABLE `users` MODIFY COLUMN `id` varchar(255) NOT NULL;

-- Modify column types in child tables
ALTER TABLE `external_service_urls` MODIFY COLUMN `user_id` varchar(255) NULL;
ALTER TABLE `skills` MODIFY COLUMN `user_id` varchar(255) NULL;

-- Step 3: Add unique index on users table if it doesn't exist
SET @index_exists = (SELECT COUNT(*) FROM INFORMATION_SCHEMA.STATISTICS 
                     WHERE TABLE_SCHEMA = DATABASE() 
                     AND TABLE_NAME = 'users' 
                     AND INDEX_NAME = 'uni_users_id');

SET @sql = IF(@index_exists = 0, 
              'ALTER TABLE `users` ADD UNIQUE INDEX `uni_users_id` (`id`)', 
              'SELECT 1');
PREPARE stmt FROM @sql;
EXECUTE stmt;
DEALLOCATE PREPARE stmt;

-- Step 4: Recreate the foreign key constraints
ALTER TABLE `external_service_urls` ADD CONSTRAINT `fk_users_external_service_urls` 
  FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON UPDATE CASCADE ON DELETE CASCADE;
  
ALTER TABLE `skills` ADD CONSTRAINT `fk_users_skills` 
  FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON UPDATE CASCADE ON DELETE CASCADE;
