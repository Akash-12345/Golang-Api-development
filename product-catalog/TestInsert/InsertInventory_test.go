package TestInsert

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/akash-searce/product-catalog/typedefs"
)

func TestInsertInventory(t *testing.T) {

	data := []byte(`{"product_id":2, "quantity":4}`)

	resp, err := http.Post("http://localhost:8089/addinventory", "application/json", bytes.NewBuffer(data))
	if err != nil {
		t.Errorf("Error making request: %v", err)
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		t.Errorf("Error reading response body: %v", err)
	}

	response := typedefs.JResponse{}
	err = json.Unmarshal(body, &response)

	if resp.StatusCode != 200 {
		t.Errorf("Expected status code 200, got %d", resp.StatusCode)
	}
	if response.Message != "Inventory detail has been added successfully" {
		t.Errorf("Expected Inventory detail has been added successfully, got %s", response.Message)
	}

}
