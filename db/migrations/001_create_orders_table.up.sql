CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE orders (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    item TEXT NOT NULL,
    customer TEXT NOT NULL,
    purchaseDate TIMESTAMP NOT NULL,
    price INT NOT NULL,
    createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
