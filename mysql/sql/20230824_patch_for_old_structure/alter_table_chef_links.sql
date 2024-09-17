ALTER TABLE `chef_links` TO `user_links`;
ALTER TABLE `user_links` RENAME COLUMN `chef_id` TO `user_id`;
ALTER TABLE `user_links` DROP FOREIGN KEY `chef_links_ibfk_1`;
ALTER TABLE `user_links` ADD CONSTRAINT `user_links_ibfk_1` FOREIGN KEY (user_id) REFERENCES users(id);