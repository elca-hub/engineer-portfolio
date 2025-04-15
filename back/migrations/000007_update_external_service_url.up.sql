ALTER TABLE external_service_urls
ADD COLUMN service_type VARCHAR(100) NOT NULL;

ALTER TABLE external_service_urls
DROP COLUMN name;
