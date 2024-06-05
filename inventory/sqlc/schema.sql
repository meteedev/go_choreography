CREATE TABLE inventory(
    id SERIAL PRIMARY KEY,
    product_code TEXT UNIQUE NOT NULL,
    product_name TEXT NOT NULL,
    description TEXT,
    quantity_in_stock INTEGER NOT NULL DEFAULT 0,
    unit_price NUMERIC NOT NULL,
    reorder_level INTEGER NOT NULL DEFAULT 0,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE
);

CREATE TABLE reservations (
    id SERIAL PRIMARY KEY,
    order_id INTEGER NOT NULL,
    product_code TEXT NOT NULL,
    quantity INTEGER NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Example insertion for a product
INSERT INTO inventory (product_code, product_name, description, quantity_in_stock, unit_price, reorder_level)
VALUES ('PROD001', 'Product 1', 'Description of Product 1', 10, 10, 20);

INSERT INTO inventory (product_code, product_name, description, quantity_in_stock, unit_price, reorder_level)
VALUES ('PROD002', 'Product 2', 'Description of Product 2', 5, 100, 20);
