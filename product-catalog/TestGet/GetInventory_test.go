package TestGet

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/akash-searce/product-catalog/typedefs"
)

func TestGetInventory(t *testing.T) {

	resp, err := http.Get("http://localhost:8089/inventorydetail/4")
	if err != nil {
		t.Errorf("Error making request: %v", err)
	}

	_, err = ioutil.ReadAll(resp.Body)

	if err != nil {
		t.Errorf("Error reading response body: %v", err)
	}

	if resp.StatusCode != 200 {
		t.Errorf("Expected status code 200, got %d", resp.StatusCode)
	}
}

func TestGetInventoryNotExist(t *testing.T) {
	resp, err := http.Get("http://localhost:8089/inventorydetail/222")
	if err != nil {
		t.Errorf("Error making request: %v", err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	response := typedefs.JResponse{}

	err = json.Unmarshal(body, &response)

	if err != nil {
		t.Errorf("Error reading response body: %v", err)
	}

	if resp.StatusCode != 200 {
		t.Errorf("Expected status code 200, got %d", resp.StatusCode)
	}

	if response.Message != "Product does not exist" {
		t.Errorf("Expected Product does not exist, got %s", response.Message)
	}

}
