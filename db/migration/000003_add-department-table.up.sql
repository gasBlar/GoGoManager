CREATE TABLE department (
    id INT PRIMARY KEY PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(33) NOT NULL,
    profileId INT,
    FOREIGN KEY (profileId) REFERENCES profileManager(id)
);