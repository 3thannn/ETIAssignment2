DROP DATABASE IF EXISTS ETIAssignment2Comment;

CREATE DATABASE ETIAssignment2Comment;

USE ETIAssignment2Comment;

DROP TABLE IF EXISTS Comment;
CREATE TABLE Comment(
	CommentID int NOT NULL PRIMARY KEY AUTO_INCREMENT,
	CreatorType varchar(255) NOT NULL,
	CreatorID int NOT NULL,
	TargetType varchar(255) NOT NULL,
	TargetID int NOT NULL,	
	CommentData varchar(500) NOT NULL,
	Anonymous bool NOT NULL DEFAULT False,
	DateTimePublished datetime);