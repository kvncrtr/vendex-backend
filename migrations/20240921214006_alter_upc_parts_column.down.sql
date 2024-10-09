-- +migrate Down
-- SQL section 'Down' is applied when this migration file is executed.

-- Reapply constraint to upc column
ALTER TABLE parts
ALTER COLUMN upc
ADD CONSTRAINT UNIQUE;

-- migrate -database "postgres://postgres:password@localhost:5432/vendex?sslmode=disable" -path ./migrations down