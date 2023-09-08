ALTER TABLE `chef_follows` TO `user_follows`;
ALTER TABLE `user_follows` DROP FOREIGN KEY `chef_follows_ibfk_1`;
ALTER TABLE `user_follows` DROP FOREIGN KEY `chef_follows_ibfk_2`;
ALTER TABLE `user_follows` ADD CONSTRAINT `user_follows_ibfk_1` FOREIGN KEY (user_id) REFERENCES users(id);
ALTER TABLE `user_follows` ADD CONSTRAINT `user_follows_ibfk_2` FOREIGN KEY (chef_id) REFERENCES users(id);