package datev_api_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestAccountingExtfJobStatus(t *testing.T) {
	req := client.NewAccountingExtfJobStatusRequest()
	req.PathParams().GUID = "867d5bed-e66a-454f-84d1-745c184ade28"

	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
