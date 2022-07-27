INSERT INTO employee VALUES(100,'David','Wallace','1967-11-17','M',NULL,NULL);

INSERT INTO branch VALUES(1,'CORPORATE',100,'2006-06-09');

UPDATE employee
SET branch_id =1
WHERE emp_id=100;

INSERT INTO employee VALUES (101,'Faizan','Ahmad','2022-07-04','M',110000,100,1);