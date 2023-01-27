package TestUpdate

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"
)

func TestUpdateProduct(t *testing.T) {
	product_master := map[string]any{
		"product_id": 1,
		"name":       "redmi",
		"specification": map[string]any{
			"key":   "processor",
			"value": "mediatek",
		},
		"sku":         "34a",
		"category_id": 1,
		"price":       2000,
	}
	product_master = map[string]any{"product_id": 2, "name": "redmi", "price": 1090}
	CheckUpdateProduct(product_master, "{\"message\":\"Product detail has been updated \"}\n", t)

}

func CheckUpdateProduct(product_master map[string]any, response string, t *testing.T) {
	json_product, err := json.Marshal(product_master)
	if err != nil {
		fmt.Println("error", err)
	}
	request_body := bytes.NewBuffer(json_product)
	req, err := http.NewRequest("PUT", "http://localhost:8089/updateproduct", request_body)
	if err != nil {
		fmt.Println("error", err)
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("error", err)
	}
	bodyBytes, err := io.ReadAll(res.Body)

	if string(bodyBytes) != response {
		t.Errorf("Expected: %v, Got: %v", response, string(bodyBytes))
	}
}
