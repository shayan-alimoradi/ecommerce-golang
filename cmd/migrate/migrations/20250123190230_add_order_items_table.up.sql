CREATE TABLE IF NOT EXISTS orders (
    id bigserial PRIMARY KEY,
    order_id int NOT NULL,
    product_id int NOT NULL,
    price decimal(10, 2) NOT NULL,
    quantity int NOT NULL,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),

    FOREIGN KEY (product_id) REFERENCES products (id) ON DELETE CASCADE,
    FOREIGN KEY (order_id) REFERENCES orders (id) ON DELETE CASCADE
);