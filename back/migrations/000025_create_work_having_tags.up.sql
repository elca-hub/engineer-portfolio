CREATE TABLE work_having_tags (
  id VARCHAR(255) NOT NULL PRIMARY KEY,
  work_id VARCHAR(255) NOT NULL,
  work_tag_id VARCHAR(255) NOT NULL,
  created_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
  updated_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
  deleted_at DATETIME(3) DEFAULT NULL,
  INDEX idx_work_having_tags_deleted_at (deleted_at),
  INDEX idx_work_having_tags_work_id (work_id),
  INDEX idx_work_having_tags_work_tag_id (work_tag_id),
  UNIQUE KEY unique_work_tag (work_id, work_tag_id),
  FOREIGN KEY (work_id) REFERENCES works(id),
  FOREIGN KEY (work_tag_id) REFERENCES work_tags(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;