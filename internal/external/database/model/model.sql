CREATE DATABASE IF NOT EXISTS cut_url;


CREATE TABLE IF NOT EXISTS users (
    id INT PRIMARY KEY NOT NULL AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    email VARCHAR (255) NOT NULL UNIQUE,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    active TINYINT(1) NOT NULL DEFAULT 1
);

CREATE TABLE IF NOT EXISTS urls (
    id INT PRIMARY KEY NOT NULL AUTO_INCREMENT,
    original VARCHAR(255) NOT NULL,
    destination VARCHAR (255) NOT NULL UNIQUE,
    user_id INT NOT NULL,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    active TINYINT(1) NOT NULL DEFAULT 1,

    CONSTRAINT fk_user
    FOREIGN KEY (user_id)
    REFERENCES users(id)
        ON DELETE CASCADE
        ON UPDATE CASCADE
);

ALTER TABLE urls ADD COLUMN description VARCHAR(100) NOT NULL;

CREATE INDEX idx_email
    ON users(email);

# CREATE INDEX idx_destination
#     ON urls(destination);
