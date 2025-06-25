ALTER TABLE works
    ADD COLUMN content_url TEXT;

ALTER TABLE works
    RENAME COLUMN thumbnail_name TO thumbnail_url;

ALTER TABLE works
    ALTER COLUMN thumbnail_url SET NOT NULL,
    ALTER COLUMN thumbnail_url SET DEFAULT '';