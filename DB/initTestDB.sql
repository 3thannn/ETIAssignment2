DROP DATABASE IF EXISTS ETIAssignment2TestDB;

CREATE DATABASE ETIAssignment2TestDB;

USE ETIAssignment2TestDB;

DROP TABLE IF EXISTS Student;
CREATE TABLE Student(
	StudentID varchar(9) NOT NULL PRIMARY KEY,
	Name varchar(255) NOT NULL);
    
INSERT INTO Student (StudentID, Name) VALUES (1, 'Ethan');
INSERT INTO Student (StudentID, Name) VALUES (2, 'Kester');
INSERT INTO Student (StudentID, Name) VALUES (3, 'Naomi');
INSERT INTO Student (StudentID, Name) VALUES (4, 'WenWei');
    
DROP TABLE IF EXISTS Tutor;
CREATE TABLE Tutor(
	TutorID varchar(9) NOT NULL PRIMARY KEY,
	Name varchar(255) NOT NULL);
INSERT INTO Tutor (TutorID, Name) VALUES (1, 'Ethan');
INSERT INTO Tutor (TutorID, Name) VALUES (2, 'Kester');
INSERT INTO Tutor (TutorID, Name) VALUES (3, 'Naomi');
INSERT INTO Tutor (TutorID, Name) VALUES (4, 'WenWei');

DROP TABLE IF EXISTS Class;
CREATE TABLE Class(
	ClassID varchar(9) NOT NULL PRIMARY KEY,
	Name varchar(255) NOT NULL);
    

INSERT INTO Class (ClassID, Name) VALUES (1, 'T01');
INSERT INTO Class (ClassID, Name) VALUES (2, 'T02');

DROP TABLE IF EXISTS Module;
CREATE TABLE Module(
	ModuleID varchar(9) NOT NULL PRIMARY KEY,
	Name varchar(255) NOT NULL);
INSERT INTO Module (ModuleID, Name) VALUES (1, 'ETI');
INSERT INTO Module (ModuleID, Name) VALUES (2, 'ERP');