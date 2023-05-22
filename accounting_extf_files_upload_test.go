package datevapi_test

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/omniboost/go-datevapi"
)

func TestAccountingExtfFilesUpload(t *testing.T) {
	f, err := os.Open("extf.csv")
	if err != nil {
		t.Error(err)
	}

	defer f.Close()
	req := client.NewAccountingExtfFilesUploadRequest()
	req.FormParams().ExtfFile = datevapi.FormFile{
		Filename: "test.csv",
		Content:  f,
	}
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	// resp.Headers.Location

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
