package datev_api_test

import (
	"log"
	"testing"

	datev_api "github.com/omniboost/go-datev-api"
)

func TestErrorResponse(t *testing.T) {
	errResp := &datev_api.ErrorResponse{
		Status: datev_api.StringInt(400),
		Title: "Bad Request",
	}

	if errResp.Error() != "Bad Request (400)" {
		log.Fatalf("Expected 'Bad Request (400)', got '%s'", errResp.Error())
	}
}
