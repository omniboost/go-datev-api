package datevapi

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-datevapi/utils"
)

func (c *Client) NewAccountingDocumentTypesGetRequest() AccountingDocumentTypesGetRequest {
	r := AccountingDocumentTypesGetRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type AccountingDocumentTypesGetRequest struct {
	client      *Client
	queryParams *AccountingDocumentTypesGetRequestQueryParams
	pathParams  *AccountingDocumentTypesGetRequestPathParams
	method      string
	headers     http.Header
	requestBody AccountingDocumentTypesGetRequestBody
}

func (r AccountingDocumentTypesGetRequest) NewQueryParams() *AccountingDocumentTypesGetRequestQueryParams {
	return &AccountingDocumentTypesGetRequestQueryParams{}
}

type AccountingDocumentTypesGetRequestQueryParams struct{}

func (p AccountingDocumentTypesGetRequestQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *AccountingDocumentTypesGetRequest) QueryParams() *AccountingDocumentTypesGetRequestQueryParams {
	return r.queryParams
}

func (r AccountingDocumentTypesGetRequest) NewPathParams() *AccountingDocumentTypesGetRequestPathParams {
	return &AccountingDocumentTypesGetRequestPathParams{}
}

type AccountingDocumentTypesGetRequestPathParams struct{}

func (p *AccountingDocumentTypesGetRequestPathParams) Params() map[string]string {
	return map[string]string{
		"service": "accounting-documents",
	}
}

func (r *AccountingDocumentTypesGetRequest) PathParams() *AccountingDocumentTypesGetRequestPathParams {
	return r.pathParams
}

func (r *AccountingDocumentTypesGetRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *AccountingDocumentTypesGetRequest) SetMethod(method string) {
	r.method = method
}

func (r *AccountingDocumentTypesGetRequest) Method() string {
	return r.method
}

func (r AccountingDocumentTypesGetRequest) NewRequestBody() AccountingDocumentTypesGetRequestBody {
	return AccountingDocumentTypesGetRequestBody{}
}

type AccountingDocumentTypesGetRequestBody struct {
}

func (r *AccountingDocumentTypesGetRequest) RequestBody() *AccountingDocumentTypesGetRequestBody {
	return nil
}

func (r *AccountingDocumentTypesGetRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *AccountingDocumentTypesGetRequest) SetRequestBody(body AccountingDocumentTypesGetRequestBody) {
	r.requestBody = body
}

func (r *AccountingDocumentTypesGetRequest) NewResponseBody() *AccountingDocumentTypesGetResponseBody {
	return &AccountingDocumentTypesGetResponseBody{}
}

type AccountingDocumentTypesGetResponseBody DocumentTypes

func (r *AccountingDocumentTypesGetRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("/v2/clients/{{.client_id}}/document-types", r.PathParams())
	return &u
}

func (r *AccountingDocumentTypesGetRequest) Do() (AccountingDocumentTypesGetResponseBody, error) {
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
