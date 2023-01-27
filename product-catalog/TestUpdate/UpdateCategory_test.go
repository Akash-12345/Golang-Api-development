package TestUpdate

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"fmt"
)

func TestUpdateCategory(t *testing.T) {
	category := map[string]any{
		"category_id":   2,
		"category_name": "clothing",
	}

	CheckUpdateCategory(category, "{\"message\":\"Category details have been updated\"}\n", t)
}

func CheckUpdateCategory(category map[string]any, response string, t *testing.T) {
	json_product, err := json.Marshal(category)
	if err != nil {
		fmt.Println("error", err)
	}

	request_body := bytes.NewBuffer(json_product)
	req, err := http.NewRequest("PUT", "http://localhost:8089/updatecategory", request_body)
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
