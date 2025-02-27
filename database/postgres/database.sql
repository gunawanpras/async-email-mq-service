CREATE TABLE users (
    id UUID PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    role_id UUID DEFAULT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    created_by UUID DEFAULT '550e8400-e29b-41d4-a716-446655440000',
    updated_at TIMESTAMP DEFAULT NULL,
    updated_by UUID DEFAULT NULL    
);

CREATE INDEX idx_users_email ON users(email);

INSERT INTO users (id, name, email, password) VALUES
('550e8400-e29b-41d4-a716-446655440000', 'SYSTEM', 'system@mq.com', '!@#$1234'),
('550e8400-e29b-41d4-a716-446655440001', 'GUNAWAN', 'gunawan@mq.com', '!@#$1234'),

CREATE TABLE email_tasks (
    id UUID PRIMARY KEY,
    sender_email VARCHAR(255) NOT NULL,
    recipient_email VARCHAR(255) NOT NULL,
    subject VARCHAR(255) NOT NULL,
    body TEXT NOT NULL,
    status VARCHAR(20) NOT NULL DEFAULT 'pending', -- possible values: pending, processing, sent, failed
    retry_count INTEGER NOT NULL DEFAULT 0,    
    created_at TIMESTAMP DEFAULT NOW(),
    created_by UUID DEFAULT NOT NULL,
    updated_at TIMESTAMP DEFAULT NULL,
    updated_by UUID DEFAULT NULL,
    CONSTRAINT fk_users FOREIGN KEY (sender_email) REFERENCES users (id) ON DELETE SET NULL
);