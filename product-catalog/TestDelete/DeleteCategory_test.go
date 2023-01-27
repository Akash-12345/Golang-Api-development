package TestDelete

import (
	"io"
	"net/http"
	"testing"

	"github.com/akash-searce/product-catalog/typedefs"
)

func TestDeleteCategoryNotExists(t *testing.T) {
	req, err := http.NewRequest("DELETE", "http://localhost:8089/deletecategory/300", nil)
	if err != nil {
		t.Fatal(err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("unexpected status code: got %d, want %d", resp.StatusCode, http.StatusOK)
	}

	expected := "{\"message\":\"The value category_id does not exist,enter a Valid ID\"}\n"

	bodyBytes, err := io.ReadAll(resp.Body)

	if string(bodyBytes) != expected {
		t.Errorf("unexpected: got %s, want %s", string(bodyBytes), expected)
	}

}

func TestDeleteCategoryExists(t *testing.T) {
	req, err := http.NewRequest("DELETE", "http://localhost:8089/deletecategory/15", nil)
	if err != nil {
		t.Fatal(err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("unexpected status code: got %d, want %d", resp.StatusCode, http.StatusOK)
	}

	response := typedefs.JResponse{}
	response.Message = "The category detail has been deleted successfully"

	bodyBytes, err := io.ReadAll(resp.Body)

	if string(bodyBytes) != response.Message {
		t.Errorf("unexpected: got %s, want %s", string(bodyBytes), response.Message)
	}

}
