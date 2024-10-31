CREATE SCHEMA IF NOT EXISTS users;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp" SCHEMA users;

CREATE TABLE IF NOT EXISTS users.user
(
    user_id       UUID PRIMARY KEY      DEFAULT users.uuid_generate_v4(),
    username      VARCHAR(255) NOT NULL,
    full_name     VARCHAR(255) NOT NULL,
    email         VARCHAR(255) NOT NULL,
    password      VARCHAR      NOT NULL,
    active        BOOLEAN      NOT NULL DEFAULT TRUE,
    hidden        BOOLEAN      NOT NULL DEFAULT FALSE,
    last_login    TIMESTAMPTZ  NOT NULL,
    custom_fields JSONB                 DEFAULT '{}'::jsonb,
    created_at    TIMESTAMPTZ  NOT NULL DEFAULT current_timestamp,
    updated_at    TIMESTAMPTZ  NOT NULL DEFAULT current_timestamp
);
