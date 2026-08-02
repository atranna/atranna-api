CREATE TABLE interfaces (
    id INTEGER PRIMARY KEY,
    device_id INT NOT NULL REFERENCES devices(id),
    ip_address TEXT NOT NULL,
    mac_address TEXT NOT NULL,
    name TEXT NOT NULL,
    state TEXT NOT NULL,
    speed INTEGER NOT NULL
);