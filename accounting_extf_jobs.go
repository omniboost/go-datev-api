package datev_api

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-datev-api/utils"
)

func (c *Client) NewAccountingExtfJobsRequest() AccountingExtfJobsRequest {
	r := AccountingExtfJobsRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type AccountingExtfJobsRequest struct {
	client      *Client
	queryParams *AccountingExtfJobsRequestQueryParams
	pathParams  *AccountingExtfJobsRequestPathParams
	method      string
	headers     http.Header
	requestBody AccountingExtfJobsRequestBody
}

func (r AccountingExtfJobsRequest) NewQueryParams() *AccountingExtfJobsRequestQueryParams {
	return &AccountingExtfJobsRequestQueryParams{}
}

type AccountingExtfJobsRequestQueryParams struct{}

func (p AccountingExtfJobsRequestQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *AccountingExtfJobsRequest) QueryParams() *AccountingExtfJobsRequestQueryParams {
	return r.queryParams
}

func (r AccountingExtfJobsRequest) NewPathParams() *AccountingExtfJobsRequestPathParams {
	return &AccountingExtfJobsRequestPathParams{}
}

type AccountingExtfJobsRequestPathParams struct {
}

func (p *AccountingExtfJobsRequestPathParams) Params() map[string]string {
	return map[string]string{
		"service": "accounting-extf-files",
	}
}

func (r *AccountingExtfJobsRequest) PathParams() *AccountingExtfJobsRequestPathParams {
	return r.pathParams
}

func (r *AccountingExtfJobsRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *AccountingExtfJobsRequest) SetMethod(method string) {
	r.method = method
}

func (r *AccountingExtfJobsRequest) Method() string {
	return r.method
}

func (r AccountingExtfJobsRequest) NewRequestBody() AccountingExtfJobsRequestBody {
	return AccountingExtfJobsRequestBody{}
}

type AccountingExtfJobsRequestBody struct {
}

func (r *AccountingExtfJobsRequest) RequestBody() *AccountingExtfJobsRequestBody {
	return nil
}

func (r *AccountingExtfJobsRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *AccountingExtfJobsRequest) SetRequestBody(body AccountingExtfJobsRequestBody) {
	r.requestBody = body
}

func (r *AccountingExtfJobsRequest) NewResponseBody() *AccountingExtfJobsResponseBody {
	return &AccountingExtfJobsResponseBody{}
}

type AccountingExtfJobsResponseBody Jobs

func (r *AccountingExtfJobsRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("/v3/clients/{{.client_id}}/extf-files/jobs", r.PathParams())
	return &u
}

func (r *AccountingExtfJobsRequest) Do() (AccountingExtfJobsResponseBody, error) {
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
	if err != nil {
		return *responseBody, err
	}

	return *responseBody, err
}
