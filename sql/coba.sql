alter table customer
ADD COLUMN email VARCHAR(100),
    ADD COLUMN balance INTEGER default 0,
    ADD COLUMN rating DOUBLE DEFAULT 0.0,
    ADD COLUMN created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    ADD COLUMN birth_date DATE,
    ADD COLUMN married BOOLEAN DEFAULT false;