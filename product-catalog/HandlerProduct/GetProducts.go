package HandlerProduct

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"strconv"

	"github.com/akash-searce/product-catalog/DbConnect"
	"github.com/akash-searce/product-catalog/Helpers"
	queries "github.com/akash-searce/product-catalog/Queries"
	"github.com/akash-searce/product-catalog/Response"
	response "github.com/akash-searce/product-catalog/Response"
	"github.com/akash-searce/product-catalog/typedefs"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
	fmt.Println("getproducts")
	pageno := r.URL.Query().Get("page")
	noofitems := r.URL.Query().Get("items_per_page")
	validpage, _ := strconv.Atoi(pageno)
	validitems, _ := strconv.Atoi(noofitems)
	if validpage < 0 {
		Helpers.SendJResponse(response.EnterValidInput, w)
		return
	}
	if validitems < 0 {
		Helpers.SendJResponse(response.EnterValidInput, w)
		return
	}

	a, _ := strconv.Atoi(pageno)
	b, _ := strconv.Atoi(noofitems)

	var ls int
	if a <= 0 && b <= 0 {
		ls = GetProductsNo(1, 20)
		b = 20
	} else if a >= 0 && b <= 0 {
		ls = GetProductsNo(a, 20)
		b = 20
	} else if a <= 0 && b >= 0 {
		ls = GetProductsNo(1, b)
	} else {
		ls = GetProductsNo(a, b)
	}

	limit_start := ls

	db := DbConnect.ConnectToDB()
	rows, err := db.Query(queries.GetProducts)
	defer rows.Close()
	if err != nil {
		Helpers.SendJResponse(response.ErrorInQuery, w)
		Helpers.HandleError(err)
		return
	}

	products := []typedefs.Product_master{}
	for rows.Next() {
		newProduct := typedefs.Product_master{}
		err = rows.Scan(&newProduct.Name, &newProduct.Price)
		products = append(products, newProduct)
	}

	response := []map[string]any{}

	for _, v := range products {
		newProduct := map[string]any{
			"name":  v.Name,
			"price": v.Price,
		}
		response = append(response, newProduct)
	}
	limit_end := int(math.Min(float64(limit_start+b), float64(len(response))))

	if limit_start <= len(response)-1 {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response[limit_start:limit_end])
		fmt.Println(response[limit_start:limit_end])
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(Response.NotEnoughProducts)
	}

}

func GetProductsNo(pn int, lpp int) int {
	limit_start := (pn - 1) * lpp

	return limit_start

}
