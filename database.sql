-- create and use sat database
CREATE DATABASE sat
USING sat

-- create users table
CREATE TABLE users (
    user_id INT PRIMARY KEY AUTO_INCREMENT,
    username VARCHAR(35) UNIQUE,
    email VARCHAR(35) UNIQUE NOT NULL,
    password VARCHAR(35) NOT NULL 
);

-- create sellers table
CREATE TABLE sellers (
    seller_id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(35) NOT NULL,
    email VARCHAR(35) UNIQUE NOT NULL 
);

-- create items table
CREATE TABLE items (
    item_id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(35) UNIQUE NOT NULL ,
    price FLOAT,
    stock INT
);

-- create orders table
CREATE TABLE orders (
    order_id INT PRIMARY KEY AUTO_INCREMENT,
    seller_id INT NOT NULL,
    amount FLOAT,
    date DATE DEFAULT NOW(),
    CONSTRAINT fk_seller_order FOREIGN KEY (seller_id)
	REFERENCES sellers(seller_id)
);

-- create order details table
CREATE TABLE order_details (
    order_detail_id INT PRIMARY KEY AUTO_INCREMENT,
    order_id INT,
    item_id INT,
    quantity INT,
    price FLOAT,
    CONSTRAINT fk__order_orderdetail FOREIGN KEY (order_id)
	REFERENCES orders(order_id),
    CONSTRAINT fk__item_orderdetail FOREIGN KEY(item_id)
    REFERENCES items(item_id)
);
