package datevapi_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestAccountingExtfJobStatus(t *testing.T) {
	req := client.NewAccountingExtfJobStatusRequest()
	req.PathParams().GUID = "cf61bb94-5ff9-4b03-831c-092602e2d8bd"

	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
