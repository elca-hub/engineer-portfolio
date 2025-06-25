CREATE TABLE certifications (
  id VARCHAR(255) NOT NULL UNIQUE PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  when_date DATETIME(3) NOT NULL,
  comment TEXT,
  sort_index INT NOT NULL,
  created_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
  updated_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
  deleted_at DATETIME(3) DEFAULT NULL,
  user_id VARCHAR(255) NOT NULL,
  INDEX idx_certifications_deleted_at (deleted_at),
  INDEX idx_certifications_user_id (user_id),
  FOREIGN KEY (user_id) REFERENCES users(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;