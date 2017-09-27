SET FOREIGN_KEY_CHECKS=0;
TRUNCATE TABLE `users`;
TRUNCATE TABLE `items`;

insert into `users` (`id`, `name`, `email`, `salt`, `salted`, `icon_image`, `created_at`, `updated_at`, `deleted_at`)
values
	(null, 'user1', 'user1@example.com', 'hoge', 'qvl6hWDUqIhwpg1W0QgVh7ClzykqMoXF5Djd525Ssf8=', '', '2017-08-30 17:55:16', '2017-08-30 17:55:16', null),
  (null, 'user2', 'user2@example.com', 'hoge', 'qvl6hWDUqIhwpg1W0QgVh7ClzykqMoXF5Djd525Ssf8=', '', '2017-08-30 17:55:16', '2017-08-30 17:55:16', null),
  (null, 'user3', 'user3@example.com', 'hoge', 'qvl6hWDUqIhwpg1W0QgVh7ClzykqMoXF5Djd525Ssf8=', '', '2017-08-30 17:55:16', '2017-08-30 17:55:16', null),
  (null, 'user4', 'user4@example.com', 'hoge', 'qvl6hWDUqIhwpg1W0QgVh7ClzykqMoXF5Djd525Ssf8=', '', '2017-08-30 17:55:16', '2017-08-30 17:55:16', null),
 	(null, 'user5', 'user5@example.com', 'hoge', 'qvl6hWDUqIhwpg1W0QgVh7ClzykqMoXF5Djd525Ssf8=', '', '2017-08-30 17:55:16', '2017-08-30 17:55:16', null),
  (null, 'user6', 'user6@example.com', 'hoge', 'qvl6hWDUqIhwpg1W0QgVh7ClzykqMoXF5Djd525Ssf8=', '', '2017-08-30 17:55:16', '2017-08-30 17:55:16', null),
  (null, 'user7', 'user7@example.com', 'hoge', 'qvl6hWDUqIhwpg1W0QgVh7ClzykqMoXF5Djd525Ssf8=', '', '2017-08-30 17:55:16', '2017-08-30 17:55:16', null),
  (null, 'user8', 'user8@example.com', 'hoge', 'qvl6hWDUqIhwpg1W0QgVh7ClzykqMoXF5Djd525Ssf8=', '', '2017-08-30 17:55:16', '2017-08-30 17:55:16', null);

insert into `items` (`id`, `user_id`, `name`, `price`, `current_payment_price`, `icon_image`, `description`, `created_at`, `updated_at`, `deleted_at`)
values
	(null, 1, 'item1', 100, 10, 'icon1', 'C', '2017-08-30 17:55:16', '2017-08-30 17:55:16', null),
  (null, 1, 'item2', 2000, 200, 'icon2', 'C#', '2017-08-30 17:55:16', '2017-08-30 17:55:16', null),
  (null, 1, 'item3', 30000, 3000, 'icon3', 'object-C', '2017-08-30 17:55:16', '2017-08-30 17:55:16', null),
  (null, 2, 'item4', 5000, 4000, 'icon4', 'Java', '2017-08-30 17:55:16', '2017-08-30 17:55:16', null),
 	(null, 2, 'item5', 50000, 5000, 'icon5', 'Javascript', '2017-08-30 17:55:16', '2017-08-30 17:55:16', null),
  (null, 3, 'item6', 500000, 5000, 'icon6', 'React', '2017-08-30 17:55:16', '2017-08-30 17:55:16', null),
  (null, 3, 'item7', 1000000, 70000, 'icon7', 'Redux', '2017-08-30 17:55:16', '2017-08-30 17:55:16', null),
  (null, 3, 'item8', 10000000, 50000, 'icon8', 'Flux', '2017-08-30 17:55:16', '2017-08-30 17:55:16', null);

