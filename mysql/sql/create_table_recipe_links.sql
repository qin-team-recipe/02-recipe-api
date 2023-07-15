CREATE TABLE IF NOT EXISTS `recipe_links` (
	`id` INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
	`recipe_id` INT(10) UNSIGNED NOT NULL,
	`url` varchar(255) NOT NULL,
	`created_at` int UNSIGNED NOT NULL,
	`updated_at` int UNSIGNED NOT NULL,
	FOREIGN KEY (recipe_id) REFERENCES recipes(id),
	PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
