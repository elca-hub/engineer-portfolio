ALTER TABLE works
    RENAME COLUMN thumbnail_url TO thumbnail_name;

ALTER TABLE works
    MODIFY COLUMN thumbnail_name VARCHAR(255) DEFAULT NULL;

