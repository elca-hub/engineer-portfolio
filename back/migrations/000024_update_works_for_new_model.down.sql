ALTER TABLE works
    DROP COLUMN content,
    DROP COLUMN github_repository_url,
    DROP COLUMN is_draft,
    ADD COLUMN github_url VARCHAR(255) NOT NULL DEFAULT '',
    ADD COLUMN thumbnail_name VARCHAR(255) DEFAULT NULL;