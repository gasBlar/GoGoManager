CREATE TABLE IF NOT EXISTS auth (
    id INT PRIMARY KEY AUTO_INCREMENT, 
    email VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL 
);