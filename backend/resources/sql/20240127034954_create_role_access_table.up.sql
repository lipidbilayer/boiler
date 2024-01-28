CREATE TABLE roles (
    name VARCHAR(25) PRIMARY KEY,
    description TEXT,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at TIMESTAMP
);

CREATE TABLE accesses (
    id SERIAL PRIMARY KEY,
    name VARCHAR(25) NOT NULL,
    menu VARCHAR(255) NOT NULL,
    type  VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at TIMESTAMP
);

CREATE TABLE role_to_accesses (
    role_name varchar(25),
    access_id INT,
    PRIMARY KEY(role_name, access_id)
);