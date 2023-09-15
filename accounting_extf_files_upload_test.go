package datevapi_test

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/omniboost/go-datevapi"
)

func TestAccountingExtfFilesUpload(t *testing.T) {
	f, err := os.Open("EXTF_schokoladenhotel_2023-04-09.csv")
	// f, err := os.Open("EXTF_C_schokoladenhotel_2023-04-09.csv")
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
		t.Fatal(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))

	time.Sleep(time.Duration(resp.RetryAfter) * time.Second)

	req2 := client.NewAccountingExtfJobStatusRequest()
	req2.PathParams().GUID = resp.GUID

	resp2, err := req2.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ = json.MarshalIndent(resp2, "", "  ")
	fmt.Println(string(b))
}
