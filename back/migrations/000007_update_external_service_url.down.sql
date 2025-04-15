ALTER TABLE external_service_urls
ADD COLUMN name VARCHAR(100) NOT NULL;

ALTER TABLE external_service_urls
DROP COLUMN service_type;
