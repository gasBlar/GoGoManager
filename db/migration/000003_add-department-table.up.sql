CREATE TABLE department (
    id INT PRIMARY KEY,
    nama VARCHAR(33) NOT NULL,
    profileId INT,
    FOREIGN KEY (profileId) REFERENCES profileManager(id)
);