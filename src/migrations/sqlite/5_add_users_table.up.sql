CREATE TABLE users (
    id INTEGER PRIMARY KEY,
    email TEXT,
    username TEXT NOT NULL UNIQUE,
    password_hash TEXT NOT NULL,
    display_name TEXT
);