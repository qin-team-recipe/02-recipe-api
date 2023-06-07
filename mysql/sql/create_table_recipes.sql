CREATE TABLE IF NOT EXISTS `recipes` (
	`id` INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
	`name` varchar(255) NOT NULL,
	`chef_id` INT(10) NOT NULL,
	`description` text NOT NULL,
	`is_draft` tinyint UNSIGNED NOT NULL,
	`created_at` int UNSIGNED NOT NULL,
	`updated_at` int UNSIGNED NOT NULL,
	`deleted_at` int UNSIGNED NULL,
	FOREIGN KEY (chef_id) REFERENCES chefs(id),
	PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
