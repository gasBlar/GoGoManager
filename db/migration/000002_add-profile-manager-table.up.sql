CREATE TABLE profileManager (
    id INT PRIMARY KEY AUTO_INCREMENT,
    authId INT NOT NULL,
    name VARCHAR(52),
    userImage VARCHAR(2083),
    companyName VARCHAR(255),
    companyImage VARCHAR(2083),
    FOREIGN KEY (authId) REFERENCES auth(id)
);