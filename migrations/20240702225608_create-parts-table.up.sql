-- +migrate Up
-- SQL section 'Up' is applied when this migration file is executed.

-- Create 'Part' table
CREATE TABLE IF NOT EXISTS parts (
   id BIGSERIAL PRIMARY KEY,
   created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
   updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
   audited_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
   part_number INTEGER UNIQUE NOT NULL DEFAULT 0,
   upc INTEGER UNIQUE DEFAULT 0,
   brand VARCHAR(100) NOT NULL DEFAULT 'vendex',   
   name VARCHAR(100) UNIQUE NOT NULL DEFAULT 'pending',  
   category VARCHAR(100) NOT NULL DEFAULT 'general', 
   description VARCHAR(2000) DEFAULT '',  
   price NUMERIC NOT NULL DEFAULT 0,
   weight NUMERIC DEFAULT 0,
   on_hand INTEGER NOT NULL,
   reorder_amount INTEGER NOT NULL,
   package_quantity INTEGER NOT NULL,
   reinventory_quantity INTEGER DEFAULT 1,
   rack_id INTEGER REFERENCES racks(id)
);

-- migrate -database "postgres://postgres:password@localhost:5432/vendex?sslmode=disable" -path ./migrations up

-- Table parts {
--   id SERIAL [pk, increment]
--   created_at TIMESTAMP [not null]
--   updated_at TIMESTAMP [not null]
--   audited_at TIMESTAMP [not null]
--   part_number BINGINT [not null]
--   upc BINGINT [not null]
--   brand VARCHAR(100) [not null]
--   name VARCHAR(100) [not null]
--   category VARCHAR(100) [not null]
--   description VARCHAR(2000) [not null]
--   price DECIMAL(2)
--   weight DECIMAL(2)
--   package_quantity INTEGER [not null]
--   on_hand INTEGER [not null]
--   reorder_amount INTEGER
--   order_id INTEGER [ref: > order.id]
--   delivery_id INTEGER [ref: > delivery.id]
--   warehouse_id INTEGER [ref: > warehouse.id]
--   inventory_id INTEGER [ref: > inventory.id]
--   transfer_id INTEGER [ref: > transfer.id]
--   rack_id INTEGER [ref: > rack.id]
-- }