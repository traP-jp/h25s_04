-- +goose Up

-- Add longitude and latitude columns to eateries table
ALTER TABLE eateries
ADD COLUMN longitude DECIMAL(10, 7) NOT NULL,
ADD COLUMN latitude DECIMAL(10, 7) NOT NULL;

-- +goose Down

-- Remove longitude and latitude columns from eateries table
ALTER TABLE eateries
DROP COLUMN longitude,
DROP COLUMN latitude;
