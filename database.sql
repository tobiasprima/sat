CREATE DATABASE sat
USING sat

CREATE TABLE sellers (
    seller_id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(35) NOT NULL,
    email VARCHAR(35) NOT NULL 
);

CREATE TABLE items (
    item_id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(35) NOT NULL ,
    price FLOAT,
    stock INT
);

CREATE TABLE orders (
    order_id INT PRIMARY KEY AUTO_INCREMENT,
    seller_id INT NOT NULL,
    amount FLOAT,
    date DATE DEFAULT NOW(),
    CONSTRAINT fk_seller_order FOREIGN KEY (seller_id)
	REFERENCES sellers(seller_id)
);

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
