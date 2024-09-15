CREATE DATABASE IF NOT EXISTS cut_url;


CREATE TABLE IF NOT EXISTS users (
    id INT PRIMARY KEY NOT NULL AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    email VARCHAR (255) NOT NULL UNIQUE,
    password VARCHAR(100) NOT NULL,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at DATETIME NULL,
    updated_at DATETIME NULL,
    active TINYINT(1) NOT NULL DEFAULT 1
);

CREATE INDEX idx_email
    ON users(email);

# CREATE INDEX idx_destination
#     ON urls(destination);
