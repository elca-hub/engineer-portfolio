CREATE TABLE works (
  id VARCHAR(255) NOT NULL PRIMARY KEY,
  title VARCHAR(255) NOT NULL,
  thumbnail_url VARCHAR(255) NOT NULL DEFAULT '',
  content_url VARCHAR(255) NOT NULL COMMENT '文章のurl',
  pinned BOOLEAN NOT NULL DEFAULT FALSE,
  sort_index INT NOT NULL DEFAULT -1 COMMENT '作品の表示順序を管理するインデックス。デフォルトは-1',
  github_url VARCHAR(255) NOT NULL DEFAULT '',
  created_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
  updated_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
  deleted_at DATETIME(3) DEFAULT NULL,
  user_id VARCHAR(255) NOT NULL,
  INDEX idx_works_deleted_at (deleted_at),
  INDEX idx_works_user_id (user_id),
  FOREIGN KEY (user_id) REFERENCES users(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;