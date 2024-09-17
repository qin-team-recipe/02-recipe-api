ALTER TABLE `recipes` ADD `user_id` INT(10) UNSIGNED NOT NULL AFTER `watch_id`;
ALTER TABLE `recipes` ADD CONSTRAINT `recipes_ibfk_1` FOREIGN KEY (user_id) REFERENCES users(id);