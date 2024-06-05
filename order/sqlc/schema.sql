CREATE TABLE orders(
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    customer_id CHARACTER(20),
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    update_reason VARCHAR(255),
    deleted_at timestamp with time zone,
    status text
);

CREATE TABLE order_items(
    id SERIAL NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    product_code text,
    unit_price numeric,
    quantity integer,
    order_id UUID,
    PRIMARY KEY(id)
);

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
