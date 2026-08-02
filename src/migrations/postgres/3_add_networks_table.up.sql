CREATE TABLE networks (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    cidr VARCHAR(45) NOT NULL,
    gateway VARCHAR(100) NOT NULL,
    vlan INT
);