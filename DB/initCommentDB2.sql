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
    
INSERT INTO Comments(CommentData, CreatorID, CreatorType, receiverId, receiverType, datetime, anonymous) 
VALUES
("Great at working in teams.", "fwna55ir8hqo57xl", "Student", "zv04w3y2tlcn5hj1", "Student", NOW(), false),
("Great team player.", "9e8uqiz7xat21opf", "Student", "zv04w3y2tlcn5hj1", "Student", NOW(), false),
("Great classmate.", "w8zuzvgadqbuift3", "Student", "zv04w3y2tlcn5hj1", "Student", NOW(), true),
("Hands up work on time.", "g8m1ce47c43blq0n", "Tutor", "zv04w3y2tlcn5hj1", "Student", NOW(), true),
("Good student.", "eg05suc92vnad01m", "Tutor", "zv04w3y2tlcn5hj1", "Student", NOW(), false),

("Awesome table partner.", "zv04w3y2tlcn5hj1", "Student", "fwna55ir8hqo57xl", "Student", NOW(), false),
("Has great leadership.", "9e8uqiz7xat21opf", "Student", "fwna55ir8hqo57xl", "Student", NOW(), false),
("Best CCA captain.", "w8zuzvgadqbuift3", "Student", "fwna55ir8hqo57xl", "Student", NOW(), true),
("Hands up work on time.", "g8m1ce47c43blq0n", "Tutor", "zv04w3y2tlcn5hj1", "Student", NOW(), true),
("Good student.", "eg05suc92vnad01m", "Tutor", "zv04w3y2tlcn5hj1", "Student", NOW(), false),

("Nice person to talk to.", "zv04w3y2tlcn5hj1", "Student", "9e8uqiz7xat21opf", "Student", NOW(), false),
("Love this person's personality.", "fwna55ir8hqo57xl", "Student", "9e8uqiz7xat21opf", "Student", NOW(), false),
("Great attitude.", "w8zuzvgadqbuift3", "Student", "9e8uqiz7xat21opf", "Student", NOW(), true),
("Hands up work on time.", "g8m1ce47c43blq0n", "Tutor", "zv04w3y2tlcn5hj1", "Student", NOW(), true),
("Good student.", "eg05suc92vnad01m", "Tutor", "zv04w3y2tlcn5hj1", "Student", NOW(), false);