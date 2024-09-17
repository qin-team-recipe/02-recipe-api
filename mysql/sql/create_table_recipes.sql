CREATE TABLE IF NOT EXISTS `recipes` (
	`id` INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
	`watch_id` varchar(50) NOT NULL,
	`user_id` INT(10) UNSIGNED NOT NULL,
	`title` varchar(255) NOT NULL,
	`description` text NULL,
	`servings` int UNSIGNED NOT NULL,
	`is_draft` tinyint UNSIGNED NOT NULL,
	`published_status` varchar(50) NOT NULL;
	`created_at` int UNSIGNED NOT NULL,
	`updated_at` int UNSIGNED NOT NULL,
	`deleted_at` int UNSIGNED NULL,
	FOREIGN KEY (user_id) REFERENCES users(id),
	PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
