ALTER TABLE `chef_images` TO `user_images`;
ALTER TABLE `user_images` RENAME COLUMN `chef_id` TO `user_id`;
ALTER TABLE `user_images` DROP FOREIGN KEY `chef_images_ibfk_1`;
ALTER TABLE `user_images` ADD CONSTRAINT `user_images_ibfk_1` FOREIGN KEY (user_id) REFERENCES users(id);