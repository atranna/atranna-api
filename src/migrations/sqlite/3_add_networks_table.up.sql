CREATE TABLE networks (
    id INTEGER PRIMARY KEY,
    name TEXT NOT NULL,
    cidr TEXT NOT NULL,
    gateway TEXT NOT NULL,
    vlan INT
);