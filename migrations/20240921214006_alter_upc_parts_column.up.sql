-- +migrate Up
-- SQL section 'Up' is applied when this migration file is executed.

-- Alter upc column
ALTER TABLE parts
ALTER COLUMN upc
DROP CONSTRAINT UNIQUE;

-- migrate -database "postgres://postgres:password@localhost:5432/vendex?sslmode=disable" -path ./migrations up