package datevapi_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestAccountingDocumentTypesGet(t *testing.T) {
	req := client.NewAccountingDocumentTypesGetRequest()

	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}

