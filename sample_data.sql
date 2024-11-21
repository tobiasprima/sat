INSERT INTO users(username, email, password) VALUES
("admin", "admin@mail.com", "12345678");


INSERT INTO sellers(name, email) VALUES
("Abraham", "abraham@mail.com"),
("Benjamin", "benjamin@mail.com"),
("Charlie", "charlie@mail.com"),
("David", "david@mail.com"),
("Edison", "edison@mail.com");

INSERT INTO items(name, price, stock) VALUES
("desktop", 11500000, 50),
("laptop", 7800000, 80),
("smart tv", 5500000, 800),
("tablet", 4200000, 100),
("smart phone", 5900000, 100),
("smart watch", 3100000, 100),
("smart lock", 2200000, 150),
("microphone", 1800000, 150),
("head set", 1100000, 200),
("laser pointer", 500000, 200);

INSERT INTO orders(seller_id, amount, date) VALUES
(1, 23000000, '2024-11-22'),
(1, 9200000, '2024-11-22'),
(1, 1800000, '2024-11-20'),
(2, 500000, '2024-11-22'),
(2, 8800000, '2024-11-15'),
(3, 22700000, '2024-11-13'),
(3, 28300000, '2024-11-22'),
(3, 5500000, '2024-11-22'),
(3, 1100000, '2024-11-13'),
(3, 1000000, '2024-11-13'),
(4, 11800000, '2024-11-22'),
(4, 57900000, '2024-11-22'),
(4, 1500000, '2024-11-22'),
(4, 11500000, '2024-11-14'),
(4, 7800000, '2024-11-14'),
(4, 5900000, '2024-11-14'),
(5, 8400000, '2024-11-20'),
(5, 8200000, '2024-11-20'),
(5, 15600000, '2024-11-18'),
(5, 500000, '2024-11-18');

INSERT INTO order_details(order_id, item_id, quantity, price) VALUES
(1, 1, 1, 11500000),
(1, 3, 2, 11000000),
(1, 10, 1, 500000),
(2, 5, 1, 5900000),
(2, 9, 3, 3300000),
(3, 8, 1, 1800000),
(4, 10, 1, 500000),
(5, 2, 1, 7800000),
(5, 10, 2, 1000000),
(6, 2, 1, 7800000),
(6, 5, 2, 11800000),
(6, 6, 1, 3100000),
(7, 1, 1, 11500000),
(7, 2, 1, 7800000),
(7, 5, 1, 5900000),
(7, 6, 1, 3100000),
(8, 3, 1, 5500000),
(9, 9, 1, 1100000),
(10, 10, 2, 1000000),
(11, 5, 1, 5900000),
(11, 5, 1, 5900000),
(12, 1, 3, 34500000),
(12, 2, 3, 23400000),
(13, 10, 1, 500000),
(13, 10, 1, 500000),
(13, 10, 1, 500000),
(14, 1, 1, 11500000),
(15, 2, 1, 7800000),
(16, 5, 1, 5900000),
(17, 4, 2, 8400000),
(18, 4, 1, 4200000),
(18, 8, 1, 1800000),
(18, 9, 2, 2200000),
(19, 2, 2, 15600000),
(20, 10, 1, 500000);
