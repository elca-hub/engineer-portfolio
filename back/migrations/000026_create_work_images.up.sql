CREATE TABLE work_images (
  id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
  file_name VARCHAR(255) NOT NULL,
  created_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
  updated_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
  deleted_at DATETIME(3) DEFAULT NULL,
  INDEX idx_work_images_deleted_at (deleted_at),
  work_id VARCHAR(255) NOT NULL,
  INDEX idx_work_images_work_id (work_id),
  FOREIGN KEY (work_id) REFERENCES works(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;