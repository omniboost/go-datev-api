package datev_api

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-datev-api/utils"
	"github.com/pkg/errors"
)

func (c *Client) NewAccountingExtfJobStatusRequest() AccountingExtfJobStatusRequest {
	r := AccountingExtfJobStatusRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type AccountingExtfJobStatusRequest struct {
	client      *Client
	queryParams *AccountingExtfJobStatusRequestQueryParams
	pathParams  *AccountingExtfJobStatusRequestPathParams
	method      string
	headers     http.Header
	requestBody AccountingExtfJobStatusRequestBody
}

func (r AccountingExtfJobStatusRequest) NewQueryParams() *AccountingExtfJobStatusRequestQueryParams {
	return &AccountingExtfJobStatusRequestQueryParams{}
}

type AccountingExtfJobStatusRequestQueryParams struct{}

func (p AccountingExtfJobStatusRequestQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *AccountingExtfJobStatusRequest) QueryParams() *AccountingExtfJobStatusRequestQueryParams {
	return r.queryParams
}

func (r AccountingExtfJobStatusRequest) NewPathParams() *AccountingExtfJobStatusRequestPathParams {
	return &AccountingExtfJobStatusRequestPathParams{}
}

type AccountingExtfJobStatusRequestPathParams struct {
	GUID string
}

func (p *AccountingExtfJobStatusRequestPathParams) Params() map[string]string {
	return map[string]string{
		"service": "accounting-extf-files",
		"guid":    p.GUID,
	}
}

func (r *AccountingExtfJobStatusRequest) PathParams() *AccountingExtfJobStatusRequestPathParams {
	return r.pathParams
}

func (r *AccountingExtfJobStatusRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *AccountingExtfJobStatusRequest) SetMethod(method string) {
	r.method = method
}

func (r *AccountingExtfJobStatusRequest) Method() string {
	return r.method
}

func (r AccountingExtfJobStatusRequest) NewRequestBody() AccountingExtfJobStatusRequestBody {
	return AccountingExtfJobStatusRequestBody{}
}

type AccountingExtfJobStatusRequestBody struct {
}

func (r *AccountingExtfJobStatusRequest) RequestBody() *AccountingExtfJobStatusRequestBody {
	return &r.requestBody
}

func (r *AccountingExtfJobStatusRequest) RequestBodyInterface() interface{} {
	return r.requestBody
}

func (r *AccountingExtfJobStatusRequest) SetRequestBody(body AccountingExtfJobStatusRequestBody) {
	r.requestBody = body
}

func (r *AccountingExtfJobStatusRequest) NewResponseBody() *AccountingExtfJobStatusResponseBody {
	return &AccountingExtfJobStatusResponseBody{}
}

type AccountingExtfJobStatusResponseBody Job

func (r *AccountingExtfJobStatusRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("/v3/clients/{{.client_id}}/extf-files/jobs/{{.guid}}", r.PathParams())
	return &u
}

func (r *AccountingExtfJobStatusRequest) Do() (AccountingExtfJobStatusResponseBody, error) {
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

	if responseBody.Result != "succeeded" {
		return *responseBody, errors.Errorf("%s: %s, affected_elements: %s", responseBody.ValidationDetails.Title, responseBody.ValidationDetails.Detail, responseBody.ValidationDetails.AffectedElements)
	}

	return *responseBody, err
}
