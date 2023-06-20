CREATE TABLE IF NOT EXISTS `user_recipes` (
	`id` INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
	`user_id` INT(10) UNSIGNED NOT NULL,
	`recipe_id` INT(10) UNSIGNED NOT NULL,
	`created_at` int UNSIGNED NOT NULL,
	`deleted_at` int UNSIGNED NULL,
	FOREIGN KEY (user_id) REFERENCES users(id),
	FOREIGN KEY (recipe_id) REFERENCES recipes(id),
	PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
