CREATE TABLE IF NOT EXISTS orders (
    id bigserial PRIMARY KEY,
    user_id int NOT NULL,
    total decimal(10, 2) NOT NULL,
    status varchar(10) NOT NULL,
    address text NOT NULL,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),

    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
);