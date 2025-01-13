CREATE TABLE department (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(33) NOT NULL,
    profileId INT,
    FOREIGN KEY (profileId) REFERENCES profileManager(id)
);
