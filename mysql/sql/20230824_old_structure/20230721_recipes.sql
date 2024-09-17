ALTER TABLE `recipes` 
	ADD `watch_id` varchar(50) NOT NULL AFTER `id`,
	-- ADD `is_limited` tinyint UNSIGNED NOT NULL AFTER `is_draft`,
	-- ADD `is_private` tinyint UNSIGNED NOT NULL AFTER `is_private`;
	ADD `published_status` varchar(50) NOT NULL AFTER `is_draft`;
