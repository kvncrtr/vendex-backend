-- +migrate Up
-- SQL section 'Up' is applied when this migration file is executed.

-- Create 'Racks' table
CREATE TABLE IF NOT EXISTS racks(
   id BIGSERIAL PRIMARY KEY,
   created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
   updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
   asset_tag VARCHAR(85) NOT NULL UNIQUE,
   work_order_id BIGSERIAL NOT NULL UNIQUE,
   is_inducted BOOLEAN NOT NULL DEFAULT FALSE,
   status VARCHAR(80) NOT NULL DEFAULT 'Waiting for induction.',
   parts VARCHAR NOT NULL,
   price NUMERIC(15, 2) NOT NULL,
   weight NUMERIC (15,2) NOT NULL,
   type VARCHAR(30) NOT NULL,
   usage VARCHAR(50) NOT NULL,
   bug_log VARCHAR DEFAULT 'No bugs associated.'
);

-- migrate -database "postgres://postgres:password@localhost:5432/vendex?sslmode=disable" -path ./migrations up