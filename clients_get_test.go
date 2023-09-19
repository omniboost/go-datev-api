package datev_api_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestClientsGet(t *testing.T) {
	req := client.NewClientsGetRequest()

	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}


