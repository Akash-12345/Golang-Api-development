package typedefs

type Product_master struct {
	Product_Id    int               `json:"product_id"`
	Name          string            `json:"name"`
	Specification map[string]string `json:"specification"` //check if it should be json
	SKU           string            `json:"sku"`
	Category_Id   int               `json:"category_id"`
	Price         float64           `json:"price"`
}

type Category_master struct {
	Category_Id   int    `json:"category_id"`
	Category_Name string `json:"category_name"`
}
type Inventory struct {
	Product_Id int `json:"product_id"`
	Quantity   int `json:"quantity"`
}

type CartItem struct {
	ProductID int `json:"product_id"`
	Quantity  int `json:"quantity"`
}
type Newcart struct {
	Price        int    `json:"price"`
	Quantity     int    `json:"quantity"`
	Product_name string `json:"product_name"`
}

type Cart struct {
	Ref       string     `json:"ref"`
	CreatedAt string     `json:"created_at"`
	Items     []CartItem `json:"items"`
	Cartvalue int        `json:"cartvalue"`
}

type JResponse struct {
	Message string `json:"message"`
}
