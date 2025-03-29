CREATE TABLE categories (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL, 
);

CREATE TABLE goods (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL, 
    price FLOAT NOT NULL,  
    category_id INT,
    FOREIGN KEY (category_id) REFERENCES categories(id) ON DELETE CASCADE,
);