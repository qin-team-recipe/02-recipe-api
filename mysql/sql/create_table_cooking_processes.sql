CREATE TABLE IF NOT EXISTS `cooking_processes` (
	`id` INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
	`recipe_id` INT(10) UNSIGNED NOT NULL,
	`title` varchar(255) NOT NULL,
	`description` text NULL,
	`step_number` INT(10) UNSIGNED NOT NULL,
	`created_at` int UNSIGNED NOT NULL,
	`updated_at` int UNSIGNED NOT NULL,
	FOREIGN KEY (recipe_id) REFERENCES recipes(id),
	PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
