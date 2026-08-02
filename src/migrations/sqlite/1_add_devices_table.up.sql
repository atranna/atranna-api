CREATE TABLE devices (
    id INTEGER PRIMARY KEY,
    hostname TEXT NOT NULL,
    ip TEXT NOT NULL,
    vendor TEXT NOT NULL,
    model TEXT NOT NULL,
    type TEXT NOT NULL
);