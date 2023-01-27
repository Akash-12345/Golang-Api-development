package queries

// queries for
var GetCategory string = "SELECT * FROM category_master WHERE category_id=$1"
var DeleteCategory string = " DELETE FROM category_master where category_id=$1"
var AddCategory string = "INSERT INTO category_master (category_id, category_name) VALUES($1,$2);"
var UpdateCategory string = "UPDATE category_master SET category_name=$1 WHERE category_id=$2;"

var AddInventory string = "INSERT INTO inventory (product_id, quantity) VALUES($1,$2);"
var DeleteInventory string = "DELETE FROM inventory WHERE product_id=$1"
var GetInventory string = "SELECT * from inventory where product_id=$1"
var UpdateInventory string = "UPDATE inventory SET quantity=$1 WHERE product_id=$2;"

var GetProduct string = "SELECT p.product_id,p.name,p.specification,p.sku,c.category_id, p.price FROM product_master p JOIN category_master c ON p.category_id =c.category_id where product_id=$1"
var DeleteProduct string = "DELETE FROM product_master WHERE product_id=$1"
var GetProducts string = "Select name,price from product_master"
var AddProduct string = "INSERT INTO product_master (product_id, name, specification,sku,category_id, price) VALUES($1,$2,$3,$4,$5,$6);"
var UpdateProduct string = "UPDATE product_master SET name=$1,sku=$2, price=$3,specification=$4 WHERE product_id =$5;"

var GetCartReference string = "SELECT * FROM cart_reference WHERE ref=$1;"
var InsertIntoCart string = "INSERT INTO cart_item VALUES($1, $2, $3);"
var GetQuantityFromCart string = "SELECT quantity FROM cart_item WHERE ref=$1 AND product_id=$2"
var UpdateCartItem string = "UPDATE cart_item SET quantity=$1 WHERE ref=$2 AND product_id=$3"
var JoinInventoryAndProductmaster string = "SELECT p.product_id, i.quantity FROM product_master p LEFT JOIN inventory i ON p.product_id=i.product_id WHERE p.product_id=$1"
var InsertCartReference string = "INSERT INTO cart_reference VALUES($1, $2);"
var GetFromCartReference string = "SELECT * FROM cart_reference WHERE ref=$1;"
var DeleteFromCart string = "DELETE FROM cart_item WHERE ref=$1 AND product_id=$2"
var JoinProductMasterCartItem = "SELECT product_master.price,product_master.name,cart_item.quantity FROM (cart_item JOIN product_master ON cart_item.product_id = product_master.product_id) WHERE cart_item.ref=$1"
