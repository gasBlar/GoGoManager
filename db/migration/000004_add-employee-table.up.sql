CREATE TABLE employee (
    id INT PRIMARY KEY, 
    identityNumber VARCHAR(33) NOT NULL, 
    name VARCHAR(33) NOT NULL, 
    employeeImageUri VARCHAR(2083), 
    gender ENUM('Male', 'Female') NOT NULL, 
    departmentId INT,
    FOREIGN KEY (departmentId) REFERENCES department(id)
);