package TestGet

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/akash-searce/product-catalog/typedefs"
)

func TestGetCartExists(t *testing.T) {

	resp, err := http.Get("http://localhost:8089/cart/get?ref=2468bb80-74cf-4af7-8096-f3a465540fb2")
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

func TestGetCartNotExists(t *testing.T) {

	resp, err := http.Get("http://localhost:8089/cart/get?ref=2468bb80-74cf-4af7-80qwe96-f3a465540f")
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

	if response.Message != "Reference_id doesn't exists" {
		t.Errorf("Expected Reference_id doesn't exists, got %s", response.Message)
	}

}
