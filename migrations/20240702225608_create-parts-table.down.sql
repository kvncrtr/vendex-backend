-- +migrate Down
-- SQL section 'Down' is applied when this migration file is executed.

-- Deletes 'Parts' table from database
DROP TABLE IF EXISTS parts;