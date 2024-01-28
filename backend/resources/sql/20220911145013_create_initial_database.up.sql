CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(25) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    name  VARCHAR(255),
    location_id INT,
    role VARCHAR(25) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    last_login_at TIMESTAMP,
    deleted_at TIMESTAMP
);

INSERT into users (username, password, name, role) 
     VALUES ('admin' ,crypt('password', gen_salt('bf')), 'admin', 'admin');
