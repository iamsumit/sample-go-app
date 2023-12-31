CREATE TABLE IF NOT EXISTS users (
    id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(128) NOT NULL,
    email VARCHAR(128) DEFAULT '',
    is_active BOOLEAN DEFAULT true,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at DATETIME
) engine = innodb;

CREATE TABLE IF NOT EXISTS user_settings (
    id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    uid INT UNSIGNED NOT NULL,
    is_subscribed BOOLEAN DEFAULT true,
    biography TEXT,
    date_of_birth DATE,
    FOREIGN KEY (uid) REFERENCES users(id)
) engine = innodb;
