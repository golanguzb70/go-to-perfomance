CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name varchar NOT NULL,
    email varchar NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

