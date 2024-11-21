## ERD Title: Smart Appliances and Tech.

1. Entities and Their Attributes:
Entity: Table_Name (e.g., Customers)

## Attributes:

A. Entity: users
- Attributes:
- user_id INT (PK, AI)
- username VARCHAR(35) UNIQUE
- email VARCHAR(35) UNIQUE NOT NULL
- password VARCHAR(35) NOT NULL

B. Entity: sellers

- Attributes:
- seller_id INT (PK, AI)
- name VARCHAR(35) NOT NULL 
- email VARCHAR(35) NOT NULL 

C. Entity: items

- Attributes:
- item_id INT (PK, AI)
- name VARCHAR(35) NOT NULL 
- price FLOAT
- stock INT

D. Entity: orders

- Attributes:
- order_id INT (PK, AI)
- seller_id INT (FK) NOT NULL
- amount FLOAT
- date DATE DEFAULT NOW

E. Entity: order_details

- Attributes:
- order_detail_id INT (PK, AI)
- order_id INT (FK)
- item_id INT (FK)
- quantity INT
- price FLOAT

## Relationships:
- Table_Name to Table_Name: (e.g., Customers to Orders)

A. Type: One to Many
- Description: one seller can have multiple orders but each order can only be done by one seller.
- sellers to orders

B. Type: One to Many
- Description: One order can have multiple order details, but each order details is linked to only one order.
- orders to order_details

C. Type: One to Many
- Description: One item can appear in many order details, but each order detail is linked to only one item.
- items to order_details

## Integrity Constraints:
- The Price in items, orders, order_details should be a positive float.
- The stock in items should be a positive integer.
- The quantity in order details should be a positive integer.
- The email in sellers should have email format.

## Additional Notes:
- The Order Details table allows the system to handle orders with multiple items.
- 