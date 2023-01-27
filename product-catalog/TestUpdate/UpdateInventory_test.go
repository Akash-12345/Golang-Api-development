package TestUpdate

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"
)

func TestUpdateInventory(t *testing.T) {
	inventory := map[string]any{
		"product_id": 29,
		"quantity":   12,
	}

	CheckUpdateInventory(inventory, "{\"message\":\"Inventory id doesn't exist\"}\n", t)
}

func CheckUpdateInventory(inventory map[string]any, response string, t *testing.T) {
	json_product, err := json.Marshal(inventory)
	if err != nil {
		fmt.Println("error", err)
	}
	request_body := bytes.NewBuffer(json_product)
	req, err := http.NewRequest("PUT", "http://localhost:8089/updateinventory", request_body)
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
