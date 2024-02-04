-- SOAL "Buatlah query SQL yang menghasilkan data sebagai berikut:"
-- id	name	parent_name
-- 1	Zaki	Ilham
-- 2	Ilham	NULL
-- 3	Irwan	Ilham
-- 4	Arka	Irwan


-- Create Database 
CREATE DATABASE IF NOT EXISTS `testdb`;

-- Create Table 
CREATE TABLE list_name (
    id INT PRIMARY KEY,
    name VARCHAR(255),
    parent_id INT,
    FOREIGN KEY (parent_id) REFERENCES list_name(id)
);

-- Insert Data to table
INSERT INTO your_table_name (id, name) VALUES
(1, 'Zaki'),
(2, 'Ilham'),
(3, 'Irwan'),
(4, 'Arka'); 

-- Update Parent_id
UPDATE list_name SET parent_id = 2 WHERE id = 1;
UPDATE list_name SET parent_id = 2 WHERE id = 3;
UPDATE list_name SET parent_id = 3 WHERE id = 4;


-- Query Jawaban Untuk Soal no 2 
SELECT t1.id, t1.name, t2.name AS parent_name FROM list_name t1 LEFT JOIN list_name t2 ON t1.parent_id = t2.id;