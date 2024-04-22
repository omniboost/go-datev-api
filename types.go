package datev_api

import (
	"time"
)

type Jobs []Job

type Job struct {
	ID                           string    `json:"id"`
	Filename                     string    `json:"filename"`
	FiscalYearBegin              Date      `json:"fiscal_year_begin"`
	ClientApplicationDisplayName string    `json:"client_application_display_name"`
	ClientApplicationVendor      string    `json:"client_application_vendor"`
	DataCategoryID               int       `json:"data_category_id"`
	DateFrom                     Date      `json:"date_from"`
	DateTo                       Date      `json:"date_to"`
	Label                        string    `json:"label"`
	NumberOfAccountingRecords    int       `json:"number_of_accounting_records"`
	Result                       string    `json:"result"`
	Timestamp                    time.Time `json:"timestamp"`
	ValidationDetails            struct {
		Type             string `json:"type"`
		Title            string `json:"title"`
		Detail           string `json:"detail"`
		AffectedElements []struct {
			Name   string `json:"name"`
			Reason string `json:"reason"`
		} `json:"affected_elements"`
	} `json:"validation_details"`
}

type DocumentTypes []DocumentType

type DocumentType struct {
	Name                  string `json:"name"`
	Category              string `json:"category"`
	DebitCreditIdentifier string `json:"debitCreditIdentifier,omitempty"`
}
