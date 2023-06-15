CREATE TABLE IF NOT EXISTS `chef_images` (
	`id` INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
	`chef_id` INT(10) UNSIGNED NOT NULL,
	`file_key` varchar(255) NULL,
	`file_name` varchar(255) NOT NULL,
	`created_at` int UNSIGNED NOT NULL,
	FOREIGN KEY (chef_id) REFERENCES chefs(id),
	PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
