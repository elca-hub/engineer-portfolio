ALTER TABLE works
    ADD COLUMN content TEXT NOT NULL,
    ADD COLUMN github_repository_url VARCHAR(255) NOT NULL DEFAULT '',
    ADD COLUMN is_draft BOOLEAN NOT NULL DEFAULT TRUE,
    DROP COLUMN github_url,
    DROP COLUMN thumbnail_name;