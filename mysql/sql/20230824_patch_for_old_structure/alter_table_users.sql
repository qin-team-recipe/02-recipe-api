ALTER TABLE `users`
    ADD `role` varchar(20) NOT NULL AFTER `email`,
    ADD `description` text NULL AFTER `role`;