CREATE TABLE IF NOT EXISTS products (
    id bigserial PRIMARY KEY,
    name varchar(255) NOT NULL,
    description TEXT NOT NULL,
    image varchar(255) NOT NULL,
    price decimal(10, 2) NOT NULL,
    quantity int NOT NULL,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW()
);