
# Product Catalog and Shopping Cart Service API

This project is to create an API which would ease the process of shopping cart service using CRUD operations.
The API would return the cart details of a customer based on a reference ID which is generated using google uuid package which will be used to view items present inside the cart.



## Getting started
### Prerequisites

- Go 1.19.3 (should still be backwards compatible with earlier versions)
- Postman-The collections will need an environment setup with `scheme`, `port` and `host` variables setup with values of `http`, `8080` and `localhost` respectively.
- PostgreSql database
## Environment Setup

Command to start working with postgresql

```bash
sudo service postgresql-9.3 initdb  (Initialize the server by running the command)
```
```bash
sudo service postgresql-9.3 start (Start the server by running the command)
```

Command to run golang code
```bash
go run main.go (Run the code using the following command)
```
Commands to upload code to github

```bash
git init -b main (Initialize the local directory as a Git repository)
```
```bash
git add . && git commit -m "initial commit"(Stage and commit all the files in your project)
```
## API Reference

### Product_master
#### Request Body
##### JSON:
- `product_id` (int): ID of the specific product.
- `name` (string): The name of the product.
- `specification` (JSON): The Specifications of the product.
- `sku` (string): Stock Keeping Unit number of the product.
- `category_id` (int): The product's category ID which needs to be present in the Category Table.
- `price` (float): The price of the product.

#### Insert Product
##### Adds a product to database

```http
  POST /addproduct
```

#### Get Product
##### Recieve product details using product id

```http
  GET /getproducts/{id:[0-9]+}
```
| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `int` | **Required**. Id of item to fetch |



#### Get Products and pagination
##### Returns a list of products based on the required pagination value. 

```http
  GET /getproducts/{page_num:[0-9]+}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `int` | **Required**. Id of item to fetch |


#### Update Product
##### Updates product detail of item depending on product_id
```http
  PUT /updateproduct
```
#### Delete Product
##### Delete product details based on product id.

```http
  Delete /deleteproduct/{id:[0-9]+}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `int` | **Required**. Id of item to fetch |


### Category_master
#### Request Body
##### JSON:
- `category_id` (int): category ID of the specific product.
- `category_name` (string): Category name of the corresponding product

#### Insert Category Name

```http
  POST /addcategory
```

#### Get Category detail

```http
  GET /getcategory/{id:[0-9]+}
```
| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `int` | **Required**. Id of item to fetch |



#### Update Category details

```http
  PUT /updatecategory
```
#### Delete Category

```http
  Delete /deletecategory/{id:[0-9]+}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `int` | **Required**. Id of item to fetch |

### Inventory
#### Request Body
##### JSON:
- `product_id` (int): Id of specific product.
- `quantity` (int): Quantity of the specific product present inside the inventory.

#### Insert inventory details

```http
  POST /addinventory
```

#### Get Inventory details

```http
  GET /inventorydetail/{id:[0-9]+}
```
| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `int` | **Required**. Id of item to fetch |



#### Update inventory detail

```http
  PUT /updateinventory
```
#### Delete inventory Detail

```http
  Delete /deleteinventory/{id:[0-9]+}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `int` | **Required**. Id of item to fetch |

### Cart 

#### Get cart details

```http
  GET /cart/get  (To retrieve the items present in the cart using the generated reference id)
```
| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `ref`      | `string` | ReferenceId used to retrieve data of a specific cart |

#### Generate cart reference id

```http
  POST /cart/createreference  (To produce a unique reference id for each customer)
```
#### Insert Product items to cart

```http
  POST /addtocart  (To add required products to cart)
```
| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `ref`      | `string` | ReferenceId used to retrieve data of a specific cart |
| `product_id`  | `int`   |Id required to insert the specific product into cart|
| `quantity`  | `int`   |Quantity of the required |


#### Delete product item from cart
```http
  DELETE /deleteitemfromcart  (Delete specific product using its product_id)
```
| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `ref`      | `string` | ReferenceId used to retrieve data of a specific cart |
| `product_id`  | `int`   |Id required to delete the specific product from the cart|







## List of packages
- google/uuid -The uuid package generates and inspects UUIDs based on RFC 4122 and DCE 1.1: Authentication and Security Services.
- encoding/json -Package json implements encoding and decoding of JSON as defined in RFC 7159.
- fmt-Package fmt implements formatted I/O with functions analogous to C's printf and scanf. 
- io/ioutil-Package ioutil implements some I/O utility functions.
- net/http-Package http provides HTTP client and server implementations.
- gorilla mux implements a request router and dispatcher for matching incoming requests to their respective handler.