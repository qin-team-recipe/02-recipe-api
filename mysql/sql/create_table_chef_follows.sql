CREATE TABLE IF NOT EXISTS `chef_follows` (
	`id` INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
	`user_id` INT(10) NOT NULL,
	`chef_id` INT(10) NOT NULL,
	`created_at` int UNSIGNED NOT NULL,
	`updated_at` int UNSIGNED NOT NULL,
	FOREIGN KEY (user_id) REFERENCES users(id),
	FOREIGN KEY (chef_id) REFERENCES chefs(id),
	PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
