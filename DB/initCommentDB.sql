DROP DATABASE IF EXISTS ETIAssignment2Comment;

CREATE DATABASE ETIAssignment2Comment;

USE ETIAssignment2Comment;

DROP TABLE IF EXISTS StudentComment;
CREATE TABLE StudentComment(
	CommentID int NOT NULL PRIMARY KEY AUTO_INCREMENT,
	CreatorID int NULL,
	SelectedStudentID int NOT NULL,	
	CommentData varchar(500) NOT NULL,
	Anonymous bool NOT NULL DEFAULT False,
	DateTimePublished datetime DEFAULT CURRENT_TIMESTAMP());

DROP TABLE IF EXISTS ClassComment;
CREATE TABLE ClassComment(
	CommentID int NOT NULL PRIMARY KEY AUTO_INCREMENT,
	CreatorID int NULL,
	ClassID int NOT NULL,
	CommentData varchar(500) NOT NULL,
	Anonymous bool NOT NULL DEFAULT False,
	DateTimePublished datetime DEFAULT CURRENT_TIMESTAMP());
    
DROP TABLE IF EXISTS ModuleComment;
CREATE TABLE ModuleComment(
	CommentID int NOT NULL PRIMARY KEY AUTO_INCREMENT,
	CreatorID int NULL,
	ModuleID int NOT NULL,
	CommentData varchar(500) NOT NULL,
	Anonymous bool NOT NULL DEFAULT False,
	DateTimePublished datetime DEFAULT CURRENT_TIMESTAMP());
    
DROP TABLE IF EXISTS TutorComment;
CREATE TABLE TutorComment(
	CommentID int NOT NULL PRIMARY KEY AUTO_INCREMENT,
	CreatorID int NULL,
	TutorID int NOT NULL,
	CommentData varchar(500) NOT NULL,
	Anonymous bool NOT NULL DEFAULT False,
	DateTimePublished datetime DEFAULT CURRENT_TIMESTAMP());
    
