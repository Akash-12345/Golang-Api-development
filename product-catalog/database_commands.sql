CREATE DATABASE product_catalog;

\c product_catalog

CREATE TABLE Product_master (
Product_Id INT PRIMARY KEY,
Name VARCHAR,
Specification JSON,
SKU VARCHAR,
Category_Id INT,
Price FLOAT,
FOREIGN KEY(Category_Id) REFERENCES Category_master(Category_Id)
);

CREATE TABLE Category_master (
Category_Id INT PRIMARY KEY,
Category_Name VARCHAR
);

CREATE TABLE Inventory (
Product_Id INT PRIMARY KEY,
Quantity INT,
FOREIGN KEY(Product_Id) REFERENCES product_master(Product_Id) ); 


CREATE TABLE cart_reference (
  ref varchar,
  created_at timestamp,
  PRIMARY KEY(ref)
);
   
CREATE TABLE cart_item (
  ref varchar,
  product_id int,
  quantity int,
  PRIMARY KEY(ref, product_id),
  FOREIGN KEY(product_id) REFERENCES product_master(product_id)
);
