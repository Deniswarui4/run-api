-- Fix existing transactions with invalid payment_metadata
-- This script updates empty string values to NULL for the payment_metadata column

-- Update transactions with empty payment_metadata to NULL
UPDATE transactions 
SET payment_metadata = NULL 
WHERE payment_metadata = '';

-- Verify the fix
SELECT COUNT(*) as fixed_count 
FROM transactions 
WHERE payment_metadata IS NULL;

-- Optional: Set default for future inserts (if not using pointer type)
-- ALTER TABLE transactions 
-- ALTER COLUMN payment_metadata SET DEFAULT NULL;
