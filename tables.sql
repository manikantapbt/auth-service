CREATE TABLE users (
                       id SERIAL PRIMARY KEY,
                       name VARCHAR(255) NOT NULL,
                       user_name VARCHAR(255) NOT NULL UNIQUE,
                       email VARCHAR(255) NOT NULL UNIQUE, -- Unique constraint on email
                       is_verified BOOLEAN NOT NULL DEFAULT FALSE,
                       country_code INT NOT NULL,
                       phone_number VARCHAR(20) UNIQUE, -- Unique constraint on phone_number
                       created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE user_events (
                             id SERIAL PRIMARY KEY,
                             phone_number VARCHAR NOT NULL,
                             event VARCHAR NOT NULL,
                             created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);