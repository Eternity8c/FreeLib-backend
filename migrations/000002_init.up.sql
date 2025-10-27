CREATE TABLE IF NOT EXISTS books (
    id SERIAL PRIMARY KEY,
    title VARCHAR(500) NOT NULL,
    author VARCHAR(500) NOT NULL,
    description TEXT,
    genre VARCHAR(500) NOT NULL,
    content TEXT,
    cover_url VARCHAR(500),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);