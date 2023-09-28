package datev_api

import (
	"encoding/json"
	"net/http"
	"net/url"
	"time"

	"github.com/omniboost/go-datev-api/utils"
)

func (c *Client) NewAccountingDocumentsPutRequest() AccountingDocumentsPutRequest {
	r := AccountingDocumentsPutRequest{
		client:  c,
		method:  http.MethodPut,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.formParams = r.NewFormParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type AccountingDocumentsPutRequest struct {
	client      *Client
	queryParams *AccountingDocumentsPutRequestQueryParams
	pathParams  *AccountingDocumentsPutRequestPathParams
	formParams  *AccountingDocumentsPutRequestFormParams
	method      string
	headers     http.Header
	requestBody AccountingDocumentsPutRequestBody
}

func (r AccountingDocumentsPutRequest) NewQueryParams() *AccountingDocumentsPutRequestQueryParams {
	return &AccountingDocumentsPutRequestQueryParams{}
}

type AccountingDocumentsPutRequestQueryParams struct{}

type AccountingDocumentsPutRequestFormParams struct {
	Metadata FileMetaData
	File     FormFile
}

func (p AccountingDocumentsPutRequestFormParams) IsMultiPart() bool {
	return true
}

func (p AccountingDocumentsPutRequestFormParams) Files() map[string]FormFile {
	return map[string]FormFile{
		"file": p.File,
	}
}

func (p AccountingDocumentsPutRequestFormParams) Values() url.Values {
	return url.Values{
		"metadata": func() []string {
			b, err := json.Marshal(p.Metadata)
			if err != nil {
				return []string{
					err.Error(),
				}
			}
			return []string{string(b)}
		}(),
	}
}

func (p AccountingDocumentsPutRequestQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *AccountingDocumentsPutRequest) QueryParams() *AccountingDocumentsPutRequestQueryParams {
	return r.queryParams
}

func (r *AccountingDocumentsPutRequest) FormParams() *AccountingDocumentsPutRequestFormParams {
	return r.formParams
}

func (r *AccountingDocumentsPutRequest) FormParamsInterface() Form {
	return r.formParams
}

func (r AccountingDocumentsPutRequest) NewPathParams() *AccountingDocumentsPutRequestPathParams {
	return &AccountingDocumentsPutRequestPathParams{}
}

func (r AccountingDocumentsPutRequest) NewFormParams() *AccountingDocumentsPutRequestFormParams {
	return &AccountingDocumentsPutRequestFormParams{}
}

type AccountingDocumentsPutRequestPathParams struct {
	GUID string
}

func (p *AccountingDocumentsPutRequestPathParams) Params() map[string]string {
	return map[string]string{
		"service": "accounting-documents",
		"guid":    p.GUID,
	}
}

func (r *AccountingDocumentsPutRequest) PathParams() *AccountingDocumentsPutRequestPathParams {
	return r.pathParams
}

func (r *AccountingDocumentsPutRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *AccountingDocumentsPutRequest) SetMethod(method string) {
	r.method = method
}

func (r *AccountingDocumentsPutRequest) Method() string {
	return r.method
}

func (r AccountingDocumentsPutRequest) NewRequestBody() AccountingDocumentsPutRequestBody {
	return AccountingDocumentsPutRequestBody{}
}

type AccountingDocumentsPutRequestBody struct {
}

func (r *AccountingDocumentsPutRequest) RequestBody() *AccountingDocumentsPutRequestBody {
	return nil
}

func (r *AccountingDocumentsPutRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *AccountingDocumentsPutRequest) SetRequestBody(body AccountingDocumentsPutRequestBody) {
	r.requestBody = body
}

func (r *AccountingDocumentsPutRequest) NewResponseBody() *AccountingDocumentsPutResponseBody {
	return &AccountingDocumentsPutResponseBody{}
}

type AccountingDocumentsPutResponseBody struct {
	ID    string `json:"id"`
	Files []struct {
		ID         string    `json:"id"`
		Name       string    `json:"name"`
		Size       string    `json:"size"`
		MediaType  string    `json:"media_type"`
		UploadDate time.Time `json:"upload_date"`
	} `json:"files"`
}

func (r *AccountingDocumentsPutRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("/v2/clients/{{.client_id}}/documents/{{.guid}}", r.PathParams())
	return &u
}

func (r *AccountingDocumentsPutRequest) Do() (AccountingDocumentsPutResponseBody, error) {
	// Create http request
	req, err := r.client.NewRequest(nil, r)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	// Process query parameters
	err = utils.AddQueryParamsToRequest(r.QueryParams(), req, false)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	responseBody := r.NewResponseBody()
	_, err = r.client.Do(req, responseBody)
	return *responseBody, err
}
