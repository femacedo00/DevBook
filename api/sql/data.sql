insert into users 
    (name, nick, email, password)
values
    ("User 1", "user_1", "user1@gmail.com", "$2a$10$JWJVMangO8mGVB9suAIsWOO/bqLlW1aHW.l6xndunqT0n7OKJRgkS"),
    ("User 2", "user_2", "user2@gmail.com", "$2a$10$JWJVMangO8mGVB9suAIsWOO/bqLlW1aHW.l6xndunqT0n7OKJRgkS"),
    ("User 3", "user_3", "user3@gmail.com", "$2a$10$JWJVMangO8mGVB9suAIsWOO/bqLlW1aHW.l6xndunqT0n7OKJRgkS"),
    ("User 4", "user_4", "user4@gmail.com", "$2a$10$JWJVMangO8mGVB9suAIsWOO/bqLlW1aHW.l6xndunqT0n7OKJRgkS"),
    ("User 5", "user_5", "user5@gmail.com", "$2a$10$JWJVMangO8mGVB9suAIsWOO/bqLlW1aHW.l6xndunqT0n7OKJRgkS"),
    ("User 6", "user_6", "user6@gmail.com", "$2a$10$JWJVMangO8mGVB9suAIsWOO/bqLlW1aHW.l6xndunqT0n7OKJRgkS");

insert into followers
    (user_id, follower_id)
values
    (1,2),
    (1,3),
    (1,5),
    (5,6),
    (4,1),
    (3,6),
    (6,2),
    (2,6);