-- Add the "tags" column to products schema if it does not exists
ALTER TABLE products ADD COLUMN IF NOT EXISTS tags jsonb;
