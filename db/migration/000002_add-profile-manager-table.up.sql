CREATE TABLE profileManager (
    id INT PRIMARY KEY AUTO_INCREMENT,
    authId INT,
    name VARCHAR(52) NOT NULL,
    userImage VARCHAR(2083),
    companyName VARCHAR(255) NOT NULL,
    companyImage VARCHAR(2083),
    FOREIGN KEY (authId) REFERENCES auth(id)
);