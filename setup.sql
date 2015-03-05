DROP DATABASE IF EXISTS goChatServer;

CREATE DATABASE goChatServer;

CREATE TABLE users (
	`Id` int(12) unsigned AUTO_INCREMENT,
	`UserType` tinyint(1) unsigned,
	`Name` varchar(50) NOT NULL,
	`NameCanonical` varchar(50) NOT NULL,
	`Password` varchar(64) NOT NULL,
	`Salt` varchar(64) NOT NULL,
	`LastLogin` timestamp NULL DEFAULT NULL,
	PRIMARY KEY (`Id`),
	UNIQUE KEY (`NameCanonical`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;
