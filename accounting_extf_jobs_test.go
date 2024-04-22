package datev_api_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestAccountingExtfJobs(t *testing.T) {
	req := client.NewAccountingExtfJobsRequest()
	// req.PathParams().GUID = "9a9cc5b8-52e0-40e2-b1ff-8c83dae59c24"

	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}

