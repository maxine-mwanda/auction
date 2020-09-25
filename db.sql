CREATE DATABASE Auction;
USE Auction;

CREATE TABLE user
(
    user_id     INT(10) PRIMARY KEY AUTO_INCREMENT,
    user_name  VARCHAR(30) NOT NULL UNIQUE,
    email VARCHAR (40) NOT NULL,
    salt VARCHAR(128) NOT NULL,
    password VARCHAR (128) NOT NULL,
    user_type VARCHAR (30) NOT NULL,
    created_at  DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at  DATETIME DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE session
(
    id INT (10) PRIMARY KEY AUTO_INCREMENT,
    user_id int(10),
    session VARCHAR (30) NOT NULL,
    expiry DATETIME,
    active BOOL,
    FOREIGN KEY (user_id) REFERENCES user (user_id)
     );


CREATE TABLE products
(
    product_id INT (40) PRIMARY KEY AUTO_INCREMENT,
    product_name varchar(100) not null,
    type VARCHAR (50) NOT NULL,
    availablity BOOL
);


-- id, name, type, starting_price, availability

CREATE TABLE auctions
(
    auction_id INT (10) PRIMARY KEY AUTO_INCREMENT,
    product_id int(10),
    active BOOL,
    FOREIGN KEY (product_id) REFERENCES products (product_id)
     );

CREATE TABLE bids
(
    bid_id INT (10) PRIMARY KEY AUTO_INCREMENT,
    auction_id int(10),
    user_id int(10),
    amount float,
    FOREIGN KEY (auction_id) REFERENCES auctions (auction_id),
    FOREIGN KEY (user_id) REFERENCES user (user_id)
     );

insert into user (user_name, email, salt, password, user_type) values ('maxine', 'maxine@mail.com', 'password', 'admin');
insert into products (type, product_name, availablity) values ('shoe', 'Highheel', 1), ('shirt', 'Beach Shirt', 0);
