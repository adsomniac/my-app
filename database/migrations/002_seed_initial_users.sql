-- Seed initial test users for development/testing
-- All test users use password: password123

INSERT INTO users (email, password_hash, role, is_active)
SELECT 'user@test.com', '$2a$10$mjPchIHjdc5d2hjxXohZEOyoZgUWGnAii453mgABT0sygBX1Nq4yO', 'user', true
WHERE NOT EXISTS (SELECT 1 FROM users WHERE LOWER(email) = 'user@test.com');

INSERT INTO users (email, password_hash, role, is_active)
SELECT 'admin@test.com', '$2a$10$mjPchIHjdc5d2hjxXohZEOyoZgUWGnAii453mgABT0sygBX1Nq4yO', 'admin', true
WHERE NOT EXISTS (SELECT 1 FROM users WHERE LOWER(email) = 'admin@test.com');

INSERT INTO users (email, password_hash, role, is_active)
SELECT 'superadmin@test.com', '$2a$10$mjPchIHjdc5d2hjxXohZEOyoZgUWGnAii453mgABT0sygBX1Nq4yO', 'super_admin', true
WHERE NOT EXISTS (SELECT 1 FROM users WHERE LOWER(email) = 'superadmin@test.com');
