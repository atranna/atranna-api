CREATE TABLE interfaces (
    id SERIAL PRIMARY KEY,
    device_id INT NOT NULL REFERENCES devices(id),
    ip_address VARCHAR(45) NOT NULL,
    mac_address VARCHAR(100) NOT NULL,
    name VARCHAR(100) NOT NULL,
    state VARCHAR(100) NOT NULL,
    speed INTEGER NOT NULL
);