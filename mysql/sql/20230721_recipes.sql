ALTER TABLE `recipes` 
	ADD `watch_id` varchar(50) NOT NULL AFTER `id`,
	ADD `is_limited` tinyint UNSIGNED NOT NULL AFTER `is_draft`;
