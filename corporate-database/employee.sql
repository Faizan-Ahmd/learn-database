DROP TABLE IF EXISTS employee;
CREATE TABLE employee(
     emp_id INT PRIMARY KEY,
     first_name VARCHAR(40),
     last_name VARCHAR(40),
     birth_day DATE,
     sex VARCHAR(1),
     salary INT,
     super_id INT,
     branch_id INT
);