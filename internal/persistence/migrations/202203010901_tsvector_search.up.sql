ALTER TABLE service ADD COLUMN tsvector_name_description tsvector;
UPDATE service SET tsvector_name_description = to_tsvector('english', coalesce(name, '') || ' ' || coalesce(description, ''));