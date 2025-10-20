-- Migration to update users table structure
-- Add username, first_name, last_name fields and remove name field

BEGIN;

-- Add new columns
ALTER TABLE users ADD COLUMN IF NOT EXISTS username VARCHAR;
ALTER TABLE users ADD COLUMN IF NOT EXISTS first_name VARCHAR;
ALTER TABLE users ADD COLUMN IF NOT EXISTS last_name VARCHAR;

-- Update existing records (if any)
-- Split existing name into first_name and last_name
UPDATE users 
SET 
    first_name = SPLIT_PART(name, ' ', 1),
    last_name = CASE 
        WHEN ARRAY_LENGTH(STRING_TO_ARRAY(name, ' '), 1) > 1 
        THEN SUBSTRING(name FROM LENGTH(SPLIT_PART(name, ' ', 1)) + 2)
        ELSE ''
    END,
    username = LOWER(email)
WHERE username IS NULL OR first_name IS NULL OR last_name IS NULL;

-- Make new columns NOT NULL after updating
ALTER TABLE users ALTER COLUMN username SET NOT NULL;
ALTER TABLE users ALTER COLUMN first_name SET NOT NULL;
ALTER TABLE users ALTER COLUMN last_name SET NOT NULL;

-- Add unique constraint on username
ALTER TABLE users ADD CONSTRAINT users_username_unique UNIQUE (username);

-- Drop old name column
ALTER TABLE users DROP COLUMN IF EXISTS name;

COMMIT;