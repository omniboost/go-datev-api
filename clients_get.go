package datev_api

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-datev-api/utils"
)

func (c *Client) NewClientsGetRequest() ClientsGetRequest {
	r := ClientsGetRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type ClientsGetRequest struct {
	client      *Client
	queryParams *ClientsGetRequestQueryParams
	pathParams  *ClientsGetRequestPathParams
	method      string
	headers     http.Header
	requestBody ClientsGetRequestBody
}

func (r ClientsGetRequest) NewQueryParams() *ClientsGetRequestQueryParams {
	return &ClientsGetRequestQueryParams{}
}

type ClientsGetRequestQueryParams struct{}

func (p ClientsGetRequestQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *ClientsGetRequest) QueryParams() *ClientsGetRequestQueryParams {
	return r.queryParams
}

func (r ClientsGetRequest) NewPathParams() *ClientsGetRequestPathParams {
	return &ClientsGetRequestPathParams{}
}

type ClientsGetRequestPathParams struct{}

func (p *ClientsGetRequestPathParams) Params() map[string]string {
	return map[string]string{
		"service": "accounting-documents",
	}
}

func (r *ClientsGetRequest) PathParams() *ClientsGetRequestPathParams {
	return r.pathParams
}

func (r *ClientsGetRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *ClientsGetRequest) SetMethod(method string) {
	r.method = method
}

func (r *ClientsGetRequest) Method() string {
	return r.method
}

func (r ClientsGetRequest) NewRequestBody() ClientsGetRequestBody {
	return ClientsGetRequestBody{}
}

type ClientsGetRequestBody struct {
}

func (r *ClientsGetRequest) RequestBody() *ClientsGetRequestBody {
	return &r.requestBody
}

func (r *ClientsGetRequest) RequestBodyInterface() interface{} {
	return r.requestBody
}

func (r *ClientsGetRequest) SetRequestBody(body ClientsGetRequestBody) {
	r.requestBody = body
}

func (r *ClientsGetRequest) NewResponseBody() *ClientsGetResponseBody {
	return &ClientsGetResponseBody{}
}

type ClientsGetResponseBody DocumentTypes

func (r *ClientsGetRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("/v2/clients/", r.PathParams())
	return &u
}

func (r *ClientsGetRequest) Do() (ClientsGetResponseBody, error) {
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
