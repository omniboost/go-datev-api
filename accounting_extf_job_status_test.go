package datev_api_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestAccountingExtfJobStatus(t *testing.T) {
	req := client.NewAccountingExtfJobStatusRequest()
	req.PathParams().GUID = "0877422c-b1af-4c43-aede-74cc24edff57"

	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
