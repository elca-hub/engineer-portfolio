ALTER TABLE works
    RENAME COLUMN thumbnail_url TO thumbnail_name;

ALTER TABLE works
    ALTER COLUMN thumbnail_name DROP NOT NULL,
    ALTER COLUMN thumbnail_name SET DEFAULT NULL;

ALTER TABLE works
    DROP COLUMN content_url;
