CREATE TABLE IF NOT EXISTS `shopping_momos` (
	`id` INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
	`user_id` INT(10) NOT NULL,
	`recipe_ingredient_id` INT(10) NOT NULL,
	`description` text NOT NULL,
	`is_done` tinyint UNSIGNED NOT NULL,
	`created_at` int UNSIGNED NOT NULL,
	`updated_at` int UNSIGNED NOT NULL,
	FOREIGN KEY (user_id) REFERENCES users(id),
	FOREIGN KEY (recipe_ingredient_id) REFERENCES recipe_ingredients(id),
	PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
