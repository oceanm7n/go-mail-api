-- Delete table Birthdays if it exists
DROP TABLE IF EXISTS Users;

-- Create table Birthdays
CREATE TABLE Users (
    id INT NOT NULL,
    name VARCHAR(50) NOT NULL,
    surname VARCHAR(50) NOT NULL,
    birthday DATE NOT NULL,
    email VARCHAR(50) NOT NULL
);

-- Populate table Birthdays with 2 records

INSERT INTO Users (id, name, surname, birthday, email) VALUES ('1', 'John', 'Smith', '1980-01-01', 'abc@mail.ru')
