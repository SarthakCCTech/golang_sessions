-- SQL dump generated using DBML (dbml.dbdiagram.io)
-- Database: MySQL
-- Generated at: 2026-10-30T07:50:28.711Z

CREATE TABLE `Users` (
  `user_id` int PRIMARY KEY AUTO_INCREMENT,
  `name` varchar(100) NOT NULL,
  `email` varchar(150) UNIQUE NOT NULL,
  `password` varchar(255) NOT NULL
);

CREATE TABLE `Todos` (
  `todo_id` int PRIMARY KEY AUTO_INCREMENT,
  `user_id` int NOT NULL,
  `title` varchar(200) NOT NULL,
  `description` text,
  `status` varchar(20) NOT NULL DEFAULT 'Pending',
  `priority` varchar(20) NOT NULL DEFAULT 'Medium',
  `created_at` timestamp NOT NULL,
  `updated_at` timestamp
);

ALTER TABLE `Todos` ADD FOREIGN KEY (`user_id`) REFERENCES `Users` (`user_id`);
