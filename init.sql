-- Create receipts table
CREATE TABLE IF NOT EXISTS receipts (
    id UUID PRIMARY KEY,
    retailer VARCHAR(255) NOT NULL,
    purchase_date DATE NOT NULL,
    purchase_time TIME NOT NULL,
    total NUMERIC NOT NULL,
    points_earned INTEGER NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Create items table
CREATE TABLE IF NOT EXISTS items (
    id SERIAL PRIMARY KEY,
    receipt_id INTEGER REFERENCES receipts(id) ON DELETE CASCADE,
    short_description VARCHAR(255) NOT NULL,
    price NUMERIC NOT NULL
);