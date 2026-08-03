CREATE TABLE users (
    id INTEGER PRIMARY KEY,
    email TEXT UNIQUE,
    username TEXT NOT NULL UNIQUE,
    password_hash TEXT NOT NULL,
    display_name TEXT
);