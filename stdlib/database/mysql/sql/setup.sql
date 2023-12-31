-- Drop the database if exists
DROP database IF EXISTS main;

-- Create the database
CREATE database main;

-- Select the database
USE main;

-- Create a table
CREATE TABLE person (
    id INT NOT NULL AUTO_INCREMENT,
    name VARCHAR(100),
    PRIMARY KEY(id)
);

-- Insert some records
INSERT INTO person (name) VALUES 
    ("Adam"), 
    ("Eve"), 
    ("George"), 
    ("Jack"), 
    ("Rita"), 
    ("Sarah"), 
    ("Luna");

-- Select all records from person table
SELECT * FROM person;
