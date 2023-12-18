CREATE TABLE IF NOT EXISTS app (
    id SERIAL PRIMARY KEY,
    name VARCHAR(128) NOT NULL,
    alias VARCHAR(128) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS activity (
    id SERIAL PRIMARY KEY,
    aid INT REFERENCES app(id) NOT NULL,
    entity VARCHAR(128) NOT NULL,
    operation VARCHAR(128) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,

    FOREIGN KEY (aid) REFERENCES app(id)
);