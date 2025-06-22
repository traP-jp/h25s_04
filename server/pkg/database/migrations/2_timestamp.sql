-- +goose Up

-- Add created_at and updated_at columns to eateries table
ALTER TABLE eateries
ADD COLUMN created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
ADD COLUMN updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;

-- Add created_at and updated_at columns to reviews table
ALTER TABLE reviews
ADD COLUMN created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
ADD COLUMN updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;

-- +goose Down

-- Remove created_at and updated_at columns from eateries table
ALTER TABLE eateries
DROP COLUMN created_at,
DROP COLUMN updated_at;

-- Remove created_at and updated_at columns from reviews table
ALTER TABLE reviews
DROP COLUMN created_at,
DROP COLUMN updated_at;
