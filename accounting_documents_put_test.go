package datev_api_test

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	datev_api "github.com/omniboost/go-datev-api"
)

func TestAccountingDocumentsPut(t *testing.T) {
	f, err := os.Open("test.pdf")
	// f, err := os.Open("EXTF_C_schokoladenhotel_2023-04-09.csv")
	if err != nil {
		t.Error(err)
	}

	defer f.Close()
	req := client.NewAccountingDocumentsPutRequest()
	req.PathParams().GUID = "9482975b-2172-4b43-b7cc-b07b007e2975"
	req.FormParams().File = datev_api.FormFile{
		Filename: "test.pdf",
		Content:  f,
	}
	req.FormParams().Metadata = datev_api.FileMetaData{
		Category: "outgoing_invoices",
		Folder:   "mews",
		Register: "2023",
	}

	resp, err := req.Do()
	if err != nil {
		t.Fatal(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
