-- +migrate Up
-- SQL section 'Up' is applied when this migration file is executed.

-- Alter upc column
ALTER TABLE parts
DROP CONSTRAINT IF EXISTS parts_upc_key;

-- migrate -database "postgres://postgres:password@localhost:5432/vendex?sslmode=disable" -path ./migrations up