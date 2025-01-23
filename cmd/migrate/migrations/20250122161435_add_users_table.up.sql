CREATE TABLE IF NOT EXISTS users (
    id bigserial PRIMARY KEY,
    username varchar(255) UNIQUE NOT NULL,
    email varchar(50) UNIQUE NOT NULL,
    password bytea,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW()
);