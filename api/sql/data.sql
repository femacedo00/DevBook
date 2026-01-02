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

insert into publications
    (title, content, author_id)
values
    ("Hello World", "This is my first publication here! Happy to join this network.", 1),
    ("Golang is awesome", "I am currently learning how to build APIs using Go. The performance is incredible!", 2),
    ("Morning routine", "Nothing beats a fresh cup of coffee and some coding early in the morning.", 3),
    ("Work-Life Balance", "Remember to take breaks. Your mental health is as important as your career.", 4),
    ("Database Design", "Spent the whole day optimizing my SQL queries. Hard work pays off!", 5),
    ("Networking", "It is not just about code; it is about the people you meet along the way.", 6),
    ("The Power of Open Source", "Contributing to open source projects is the best way to level up your skills.", 7),
    ("API Security", "Always remember to use JWT and bcrypt to protect your users' data.", 2);