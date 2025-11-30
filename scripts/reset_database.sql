-- Reset Database Script
-- This will delete all users and related data, then create fresh test users

-- Disable foreign key checks temporarily
SET session_replication_role = replica;

-- Delete all data in correct order (respecting foreign keys)
DELETE FROM tickets;
DELETE FROM transactions;
DELETE FROM withdrawals;
DELETE FROM ticket_types;
DELETE FROM events;
DELETE FROM organizer_balances;
DELETE FROM users;

-- Re-enable foreign key checks
SET session_replication_role = DEFAULT;

-- Insert fresh test users with hashed passwords
-- Password for all users: Admin@123, Moderator@123, Organizer@123

-- Admin User (pre-verified)
INSERT INTO users (
    id, email, password, first_name, last_name, phone, role, 
    is_active, is_verified, created_at, updated_at
) VALUES (
    gen_random_uuid(),
    'admin@eventtickets.com',
    '$2a$10$rpdnlUKtgvxWnOTVZG.XuOK.8QB9JXbxq6rODsTlbUiP1ZIRQMvHa', -- Admin@123
    'System',
    'Administrator',
    '+1234567890',
    'admin',
    true,
    true,
    NOW(),
    NOW()
);

-- Moderator User (pre-verified)
INSERT INTO users (
    id, email, password, first_name, last_name, phone, role, 
    is_active, is_verified, created_at, updated_at
) VALUES (
    gen_random_uuid(),
    'moderator@eventtickets.com',
    '$2a$10$rpdnlUKtgvxWnOTVZG.XuOK.8QB9JXbxq6rODsTlbUiP1ZIRQMvHa', -- Moderator@123
    'Test',
    'Moderator',
    '+1234567891',
    'moderator',
    true,
    true,
    NOW(),
    NOW()
);

-- Organizer User (pre-verified)
INSERT INTO users (
    id, email, password, first_name, last_name, phone, role, 
    is_active, is_verified, created_at, updated_at
) VALUES (
    gen_random_uuid(),
    'organizer@eventtickets.com',
    '$2a$10$rpdnlUKtgvxWnOTVZG.XuOK.8QB9JXbxq6rODsTlbUiP1ZIRQMvHa', -- Organizer@123
    'Test',
    'Organizer',
    '+1234567892',
    'organizer',
    true,
    true,
    NOW(),
    NOW()
);

-- Create organizer balance for the organizer
INSERT INTO organizer_balances (
    id, organizer_id, total_earnings, available_balance, 
    pending_balance, withdrawn_amount, created_at, updated_at
)
SELECT 
    gen_random_uuid(),
    id,
    0,
    0,
    0,
    0,
    NOW(),
    NOW()
FROM users 
WHERE role = 'organizer' AND email = 'organizer@eventtickets.com';

-- Show results
SELECT 
    email, 
    role, 
    is_verified, 
    is_active,
    first_name || ' ' || last_name as full_name
FROM users 
ORDER BY role;
