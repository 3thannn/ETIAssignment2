DROP DATABASE IF EXISTS ETIAssignment2Rating;

CREATE DATABASE ETIAssignment2Rating;

USE ETIAssignment2Rating;

DROP TABLE IF EXISTS StudentRating;
CREATE TABLE StudentRating(
	RatingID int NOT NULL PRIMARY KEY AUTO_INCREMENT,
	CreatorID int NULL,
	SelectedStudentID int NOT NULL,
	RatingScore int NOT NULL,
	Anonymous bool NOT NULL DEFAULT False,
	DateTimePublished datetime DEFAULT CURRENT_TIMESTAMP());

DROP TABLE IF EXISTS ClassRating;
CREATE TABLE ClassRating(
	RatingID int NOT NULL PRIMARY KEY AUTO_INCREMENT,
	CreatorID int NULL,
	ClassID int NOT NULL,
	RatingScore int NOT NULL,
	Anonymous bool NOT NULL DEFAULT False,
	DateTimePublished datetime DEFAULT CURRENT_TIMESTAMP());
    
DROP TABLE IF EXISTS ModuleRating;
CREATE TABLE ModuleRating(
	RatingID int NOT NULL PRIMARY KEY AUTO_INCREMENT,
	CreatorID int NULL,
	ModuleID int NOT NULL,
	RatingScore int NOT NULL,
	Anonymous bool NOT NULL DEFAULT False,
	DateTimePublished datetime DEFAULT CURRENT_TIMESTAMP());
    
DROP TABLE IF EXISTS TutorRating;
CREATE TABLE TutorRatinng(
	RatingID int NOT NULL PRIMARY KEY AUTO_INCREMENT,
	CreatorID int NULL,
	TutorID int NOT NULL,
	RatingScore int NOT NULL,
	Anonymous bool NOT NULL DEFAULT False,
	DateTimePublished datetime DEFAULT CURRENT_TIMESTAMP());
    
