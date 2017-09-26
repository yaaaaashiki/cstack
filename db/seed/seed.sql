SET FOREIGN_KEY_CHECKS=0;
TRUNCATE TABLE `users`;

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

