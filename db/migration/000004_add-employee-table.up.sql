CREATE TABLE employee (
    id INT AUTO_INCREMENT PRIMARY KEY, 
    identityNumber VARCHAR(33) NOT NULL, 
    name VARCHAR(33) NOT NULL, 
    employeeImageUri VARCHAR(2083), 
    gender ENUM('male', 'female') NOT NULL, 
    departmentId INT,
    FOREIGN KEY (departmentId) REFERENCES department(id)
);