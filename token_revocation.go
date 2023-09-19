package datev_api

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-datev-api/utils"
)

func (c *Client) NewTokenRevocationRequest() TokenRevocationRequest {
	r := TokenRevocationRequest{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	// r.formParams = r.NewFormParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type TokenRevocationRequest struct {
	client      *Client
	queryParams *TokenRevocationRequestQueryParams
	pathParams  *TokenRevocationRequestPathParams
	// formParams  *TokenRevocationRequestFormparams
	method      string
	headers     http.Header
	requestBody TokenRevocationRequestBody
}

func (r TokenRevocationRequest) NewQueryParams() *TokenRevocationRequestQueryParams {
	return &TokenRevocationRequestQueryParams{}
}

type TokenRevocationRequestQueryParams struct{}

func (p TokenRevocationRequestQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *TokenRevocationRequest) QueryParams() *TokenRevocationRequestQueryParams {
	return r.queryParams
}

// type TokenRevocationRequestFormparams struct {
// 	ClientID string
// 	ClientSecret string
// 	Token string
// 	TokenTypeHint string
// }

// func (p TokenRevocationRequestFormparams) Files() map[string]FormFile {
// 	return map[string]FormFile{}
// }

// func (p TokenRevocationRequestFormparams) Values() url.Values {
// 	return url.Values{
// 		"client_id": []string{p.ClientID},
// 		"client_secret": []string{p.ClientSecret},
// 		"token": []string{p.Token},
// 		"token_type_hint": []string{p.TokenTypeHint},
// 	}
// }

// func (r TokenRevocationRequest) NewFormParams() *TokenRevocationRequestFormparams {
// 	return &TokenRevocationRequestFormparams{}
// }

// func (r *TokenRevocationRequest) FormParams() *TokenRevocationRequestFormparams {
// 	return r.formParams
// }

// func (r *TokenRevocationRequest) FormParamsInterface() Form {
// 	return r.formParams
// }

func (r TokenRevocationRequest) NewPathParams() *TokenRevocationRequestPathParams {
	return &TokenRevocationRequestPathParams{}
}

type TokenRevocationRequestPathParams struct{}

func (p *TokenRevocationRequestPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *TokenRevocationRequest) PathParams() *TokenRevocationRequestPathParams {
	return r.pathParams
}

func (r *TokenRevocationRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *TokenRevocationRequest) SetMethod(method string) {
	r.method = method
}

func (r *TokenRevocationRequest) Method() string {
	return r.method
}

func (r TokenRevocationRequest) NewRequestBody() TokenRevocationRequestBody {
	return TokenRevocationRequestBody{}
}

type TokenRevocationRequestBody struct {
	Token         string `json:"token"`
	TokenTypeHint string `json:"token_type_hint,omitempty"`
}

func (r *TokenRevocationRequest) RequestBody() *TokenRevocationRequestBody {
	return &r.requestBody
}

func (r *TokenRevocationRequest) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *TokenRevocationRequest) SetRequestBody(body TokenRevocationRequestBody) {
	r.requestBody = body
}

func (r *TokenRevocationRequest) NewResponseBody() *TokenRevocationResponseBody {
	return &TokenRevocationResponseBody{}
}

type TokenRevocationResponseBody struct{}

func (r *TokenRevocationRequest) URL() *url.URL {
	u, _ := url.Parse(r.client.RevokeURL())
	return u
}

func (r *TokenRevocationRequest) Do() (TokenRevocationResponseBody, error) {
	// Create http request
	req, err := r.client.NewRequest(nil, r)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// Process query parameters
	err = utils.AddQueryParamsToRequest(r.QueryParams(), req, false)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	responseBody := r.NewResponseBody()
	_, err = r.client.Do(req, responseBody)

	return *responseBody, err
}
