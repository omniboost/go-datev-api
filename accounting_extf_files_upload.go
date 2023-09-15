package datevapi

import (
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/omniboost/go-datevapi/utils"
)

func (c *Client) NewAccountingExtfFilesUploadRequest() AccountingExtfFilesUploadRequest {
	r := AccountingExtfFilesUploadRequest{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.formParams = r.NewFormParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type AccountingExtfFilesUploadRequest struct {
	client      *Client
	queryParams *AccountingExtfFilesUploadRequestQueryParams
	pathParams  *AccountingExtfFilesUploadRequestPathParams
	formParams  *AccountingExtfFilesUploadRequestFormParams
	method      string
	headers     http.Header
	requestBody AccountingExtfFilesUploadRequestBody
}

func (r AccountingExtfFilesUploadRequest) NewQueryParams() *AccountingExtfFilesUploadRequestQueryParams {
	return &AccountingExtfFilesUploadRequestQueryParams{}
}

type AccountingExtfFilesUploadRequestQueryParams struct{}

type AccountingExtfFilesUploadRequestFormParams struct {
	ExtfFile FormFile
}

func (p AccountingExtfFilesUploadRequestFormParams) Files() map[string]FormFile {
	return map[string]FormFile{
		"extf-file": p.ExtfFile,
	}
}

func (p AccountingExtfFilesUploadRequestFormParams) Values() url.Values {
	return url.Values{}
}

func (p AccountingExtfFilesUploadRequestQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *AccountingExtfFilesUploadRequest) QueryParams() *AccountingExtfFilesUploadRequestQueryParams {
	return r.queryParams
}

func (r *AccountingExtfFilesUploadRequest) FormParams() *AccountingExtfFilesUploadRequestFormParams {
	return r.formParams
}

func (r AccountingExtfFilesUploadRequest) NewPathParams() *AccountingExtfFilesUploadRequestPathParams {
	return &AccountingExtfFilesUploadRequestPathParams{}
}

func (r AccountingExtfFilesUploadRequest) NewFormParams() *AccountingExtfFilesUploadRequestFormParams {
	return &AccountingExtfFilesUploadRequestFormParams{}
}

type AccountingExtfFilesUploadRequestPathParams struct {
	Service string
}

func (p *AccountingExtfFilesUploadRequestPathParams) Params() map[string]string {
	return map[string]string{
		"service": "accounting-extf-files",
	}
}

func (r *AccountingExtfFilesUploadRequest) PathParams() *AccountingExtfFilesUploadRequestPathParams {
	return r.pathParams
}

func (r *AccountingExtfFilesUploadRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *AccountingExtfFilesUploadRequest) SetMethod(method string) {
	r.method = method
}

func (r *AccountingExtfFilesUploadRequest) Method() string {
	return r.method
}

func (r AccountingExtfFilesUploadRequest) NewRequestBody() AccountingExtfFilesUploadRequestBody {
	return AccountingExtfFilesUploadRequestBody{}
}

type AccountingExtfFilesUploadRequestBody struct {
}

func (r *AccountingExtfFilesUploadRequest) RequestBody() *AccountingExtfFilesUploadRequestBody {
	return nil
}

func (r *AccountingExtfFilesUploadRequest) RequestBodyInterface() interface{} {
	// Convert file to io.Reader
	return r.FormParams().ExtfFile.Content
}

func (r *AccountingExtfFilesUploadRequest) SetRequestBody(body AccountingExtfFilesUploadRequestBody) {
	r.requestBody = body
}

func (r *AccountingExtfFilesUploadRequest) NewResponseBody() *AccountingExtfFilesUploadResponseBody {
	return &AccountingExtfFilesUploadResponseBody{}
}

type AccountingExtfFilesUploadResponseBody struct {
	Location   string
	GUID       string
	RetryAfter int
}

func (r *AccountingExtfFilesUploadRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("/v3/clients/{{.client_id}}/extf-files/import", r.PathParams())
	return &u
}

func (r *AccountingExtfFilesUploadRequest) Do() (AccountingExtfFilesUploadResponseBody, error) {
	// Create http request
	req, err := r.client.NewRequest(nil, r)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	req.Header.Add("Filename", "EXTF_Buchungsstapel.csv")
	// req.Header.Add("Filename", r.FormParams().ExtfFile.Filename)
	req.Header.Set("Content-Type", "application/octet-stream")

	// Process query parameters
	err = utils.AddQueryParamsToRequest(r.QueryParams(), req, false)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	responseBody := r.NewResponseBody()
	resp, err := r.client.Do(req, responseBody)
	responseBody.Location = resp.Header.Get("Location")
	responseBody.RetryAfter, _ = strconv.Atoi(resp.Header.Get("Retry-After"))
	pieces := strings.Split(responseBody.Location, "/")
	responseBody.GUID = pieces[len(pieces)-1]
	return *responseBody, err
}
