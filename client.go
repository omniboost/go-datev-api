package datev_api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"net/http/httputil"
	"net/textproto"
	"net/url"
	"path"
	"strings"
	"text/template"

	"github.com/pkg/errors"
)

const (
	libraryVersion = "0.0.1"
	userAgent      = "go-datev/" + libraryVersion
	mediaType      = "application/json"
	charset        = "utf-8"
)

var (
	BaseURL   = "https://{{.service}}.api.datev.de/platform/"
	RevokeURL = "https://api.datev.de/revoke"
)

// NewClient returns a new Exact Globe Client client
func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	client := &Client{}

	client.SetHTTPClient(httpClient)
	client.SetBaseURL(BaseURL)
	client.SetRevokeURL(RevokeURL)
	client.SetDebug(false)
	client.SetUserAgent(userAgent)
	client.SetMediaType(mediaType)
	client.SetCharset(charset)

	return client
}

// Client manages communication with Exact Globe Client
type Client struct {
	// HTTP client used to communicate with the Client.
	http *http.Client

	debug     bool
	baseURL   string
	revokeURL string

	// credentials
	clientID      string
	clientSecret  string
	datevClientID string

	oauth *Oauth2Config

	// User agent for client
	userAgent string

	mediaType             string
	charset               string
	disallowUnknownFields bool

	// Optional function called after every successful request made to the DO Clients
	beforeRequestDo    BeforeRequestDoCallback
	onRequestCompleted RequestCompletionCallback
}

type BeforeRequestDoCallback func(*http.Client, *http.Request, interface{})

// RequestCompletionCallback defines the type of the request callback function
type RequestCompletionCallback func(*http.Request, *http.Response)

func (c *Client) SetHTTPClient(client *http.Client) {
	c.http = client
}

func (c *Client) HTTPClient() *http.Client {
	return c.http
}

func (c Client) Debug() bool {
	return c.debug
}

func (c *Client) SetDebug(debug bool) {
	c.debug = debug
}

func (c Client) BaseURL() string {
	return c.baseURL
}

func (c *Client) SetBaseURL(baseURL string) {
	c.baseURL = baseURL
}

func (c Client) RevokeURL() string {
	return c.revokeURL
}

func (c *Client) SetRevokeURL(revokeURL string) {
	c.revokeURL = revokeURL
}

func (c *Client) SetMediaType(mediaType string) {
	c.mediaType = mediaType
}

func (c Client) MediaType() string {
	return mediaType
}

func (c *Client) SetCharset(charset string) {
	c.charset = charset
}

func (c Client) Charset() string {
	return charset
}

func (c *Client) SetClientID(clientID string) {
	c.clientID = clientID
}

func (c Client) ClientID() string {
	return c.clientID
}

func (c *Client) SetClientSecret(clientSecret string) {
	c.clientSecret = clientSecret
}

func (c Client) ClientSecret() string {
	return c.clientSecret
}

func (c *Client) SetDatevClientID(datevClientID string) {
	c.datevClientID = datevClientID
}

func (c Client) DatevClientID() string {
	return c.datevClientID
}

func (c *Client) SetOauth(oauth *Oauth2Config) {
	c.oauth = oauth
}

func (c Client) Oauth() *Oauth2Config {
	return c.oauth
}

func (c *Client) SetUserAgent(userAgent string) {
	c.userAgent = userAgent
}

func (c Client) UserAgent() string {
	return userAgent
}

func (c *Client) SetDisallowUnknownFields(disallowUnknownFields bool) {
	c.disallowUnknownFields = disallowUnknownFields
}

func (c *Client) SetBeforeRequestDo(fun BeforeRequestDoCallback) {
	c.beforeRequestDo = fun
}

func (c *Client) GetEndpointURL(p string, pathParams PathParams) url.URL {
	params := pathParams.Params()
	params["client_id"] = c.DatevClientID()

	baseURL := c.BaseURL()

	tmpl, err := template.New("url").Parse(baseURL)
	if err != nil {
		log.Fatal(err)
	}

	buf := new(bytes.Buffer)
	err = tmpl.Execute(buf, params)
	if err != nil {
		log.Fatal(err)
	}

	b, err := url.Parse(buf.String())
	if err != nil {
		log.Fatal(err)
	}

	tmpl, err = template.New("endpoint").Parse(p)
	if err != nil {
		log.Fatal(err)
	}

	buf = new(bytes.Buffer)
	err = tmpl.Execute(buf, params)
	if err != nil {
		log.Fatal(err)
	}

	e, err := url.Parse(buf.String())
	if err != nil {
		log.Fatal(err)
	}

	// combine the two
	q := b.Query()
	for k, vv := range e.Query() {
		for _, v := range vv {
			q.Add(k, v)
		}
	}

	b.RawQuery = q.Encode()
	b.Path = path.Join(b.Path, e.Path)

	return *b
}

func (c *Client) NewRequest(ctx context.Context, req Request) (*http.Request, error) {
	var contentType string

	// convert body struct to json
	buf := new(bytes.Buffer)

	if req.RequestBodyInterface() != nil {
		// Request has a method that returns a request body
		if r, ok := req.RequestBodyInterface().(io.Reader); ok {
			// request body is a io.Reader
			_, err := io.Copy(buf, r)
			if err != nil {
				return nil, err
			}
		} else {
			// request body is a struct/slice; marshal to json
			err := json.NewEncoder(buf).Encode(req.RequestBodyInterface())
			if err != nil {
				return nil, err
			}
		}
	} else if i, ok := req.(interface{ FormParamsInterface() Form }); ok {
		if i.FormParamsInterface().IsMultiPart() {
			// @TODO implement this as RequestBodyInterface()
			// Request has a form as body
			var err error
			w := multipart.NewWriter(buf)

			for k, f := range i.FormParamsInterface().Files() {
				var part io.Writer
				if x, ok := f.Content.(io.Closer); ok {
					defer x.Close()
				}

				if part, err = w.CreateFormFile(k, f.Filename); err != nil {
					return nil, err
				}

				if _, err = io.Copy(part, f.Content); err != nil {
					return nil, err
				}
			}

			for k := range i.FormParamsInterface().Values() {
				var part io.Writer

				// Add other fields
				if part, err = w.CreateFormField(k); err != nil {
					return nil, err
				}

				fv := strings.NewReader(i.FormParamsInterface().Values().Get(k))
				if _, err = io.Copy(part, fv); err != nil {
					return nil, err
				}
			}

			// Don't forget to close the multipart writer.
			// If you don't close it, your request will be missing the terminating boundary.
			w.Close()

			// Don't forget to set the content type, this will contain the boundary.
			contentType = w.FormDataContentType()
		} else {
			buf.WriteString(i.FormParamsInterface().Values().Encode())
		}
	}

	// create new http request
	r, err := http.NewRequest(req.Method(), req.URL().String(), buf)
	if err != nil {
		return nil, err
	}

	// values := url.Values{}
	// err = utils.AddURLValuesToRequest(values, req, true)
	// if err != nil {
	// 	return nil, err
	// }

	// optionally pass along context
	if ctx != nil {
		r = r.WithContext(ctx)
	}

	// set other headers
	if contentType != "" {
		r.Header.Add("Content-Type", contentType)
	} else {
		r.Header.Add("Content-Type", fmt.Sprintf("%s; charset=%s", c.MediaType(), c.Charset()))
	}
	r.Header.Add("Accept", c.MediaType())
	r.Header.Add("User-Agent", c.UserAgent())
	r.Header.Add("X-DATEV-Client-Id", c.ClientID())

	return r, nil
}

// Do sends an Client request and returns the Client response. The Client response is json decoded and stored in the value
// pointed to by v, or returned as an error if an Client error has occurred. If v implements the io.Writer interface,
// the raw response will be written to v, without attempting to decode it.
func (c *Client) Do(req *http.Request, body interface{}) (*http.Response, error) {
	if c.beforeRequestDo != nil {
		c.beforeRequestDo(c.http, req, body)
	}

	if c.debug == true {
		dump, _ := httputil.DumpRequestOut(req, true)
		log.Println(string(dump))
	}

	httpResp, err := c.http.Do(req)
	if err != nil {
		return nil, err
	}

	if c.onRequestCompleted != nil {
		c.onRequestCompleted(req, httpResp)
	}

	// close body io.Reader
	defer func() {
		if rerr := httpResp.Body.Close(); err == nil {
			err = rerr
		}
	}()

	if c.debug == true {
		dump, _ := httputil.DumpResponse(httpResp, true)
		log.Println(string(dump))
	}

	// check if the response isn't an error
	err = CheckResponse(httpResp)
	if err != nil {
		return httpResp, err
	}

	// check the provided interface parameter
	if httpResp == nil {
		return httpResp, nil
	}

	if body == nil {
		return httpResp, err
	}

	if httpResp.ContentLength == 0 {
		return httpResp, nil
	}

	errResp := &ErrorResponse{Response: httpResp}
	err = c.Unmarshal(httpResp.Body, body, errResp)
	if err != nil {
		return httpResp, err
	}

	if errResp.Error() != "" {
		return httpResp, errResp
	}

	return httpResp, nil
}

func (c *Client) Unmarshal(r io.Reader, vv ...interface{}) error {
	if len(vv) == 0 {
		return nil
	}

	b, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}

	errs := []error{}
	for _, v := range vv {
		r := bytes.NewReader(b)
		dec := json.NewDecoder(r)
		if c.disallowUnknownFields {
			dec.DisallowUnknownFields()
		}

		err := dec.Decode(v)
		if err != nil && err != io.EOF {
			errs = append(errs, err)
		}

	}

	if len(errs) == len(vv) {
		// Everything errored
		msgs := make([]string, len(errs))
		for i, e := range errs {
			msgs[i] = fmt.Sprint(e)
		}
		return errors.New(strings.Join(msgs, ", "))
	}

	return nil
}

// CheckResponse checks the Client response for errors, and returns them if
// present. A response is considered an error if it has a status code outside
// the 200 range. Client error responses are expected to have either no response
// body, or a json response body that maps to ErrorResponse. Any other response
// body will be silently ignored.
func CheckResponse(r *http.Response) error {
	errorResponse := &ErrorResponse{Response: r}

	// Don't check content-lenght: a created response, for example, has no body
	// if r.Header.Get("Content-Length") == "0" {
	// 	errorResponse.Errors.Message = r.Status
	// 	return errorResponse
	// }

	if c := r.StatusCode; c >= 200 && c <= 299 {
		return nil
	}

	// read data and copy it back
	data, err := ioutil.ReadAll(r.Body)
	r.Body = ioutil.NopCloser(bytes.NewReader(data))
	if err != nil {
		return errorResponse
	}

	// err = checkContentType(r)
	// if err != nil {
	// 	return errors.WithStack(err)
	// }

	if r.ContentLength == 0 {
		return errors.New("response body is empty")
	}

	// convert json to struct
	if len(data) != 0 {
		err = json.Unmarshal(data, &errorResponse)
		if err != nil {
			return errors.WithStack(err)
		}
	}

	if errorResponse.Error() != "" {
		return errorResponse
	}

	return nil
}

// {
//   "timestamp": "2023-05-12T13:29:05.028+0000",
//   "status": 404,
//   "error": "Not Found",
//   "exception": "com.datev.financial.service.exceptions.ResourceNotFoundException",
//   "message": "Business Location Not found for 12345",
//   "path": "/f/finance/12345/financials/2023-05-10T09:00:00+02:00/2023-05-12T09:00:00+02:00"
// }

// {
//   "httpCode": "401",
//   "httpMessage": "Unauthorized",
//   "moreInformation": "this offline access token is not authorized to access requested clients",
//   "status": "401",
//   "title": "Unauthorized",
//   "detail": "this offline access token is not authorized to access requested clients"
// }

type ErrorResponse struct {
	// HTTP response that caused this error
	Response *http.Response

	Title           string `json:"title"`
	RequestID       string `json:"request_id"`
	HTTPCode        string `json:"httpCode"`
	HTTPMessage     string `json:"httpMessage"`
	MoreInformation string `json:"moreInformation"`
	Status          int    `json:"status"`
	Detail          string `json:"detail"`
}

func (r *ErrorResponse) Error() string {
	if r.Title == "" {
		return ""
	}

	details := []string{}
	if r.MoreInformation != "" {
		details = append(details, r.MoreInformation)
	}
	if r.Detail != "" {
		details = append(details, r.Detail)
	}

	if len(details) > 0 {
		return fmt.Sprintf("%s (%d): %s", r.Title, r.Status, strings.Join(details, " "))
	}

	return fmt.Sprintf("%s (%d)", r.Title, r.Status)
}

func checkContentType(response *http.Response) error {
	header := response.Header.Get("Content-Type")
	contentType := strings.Split(header, ";")[0]
	if contentType != mediaType {
		return fmt.Errorf("Expected Content-Type \"%s\", got \"%s\"", mediaType, contentType)
	}

	return nil
}

func CreateFormFile(w *multipart.Writer, data io.Reader, fieldname, filename string) (io.Writer, error) {
	var quoteEscaper = strings.NewReplacer("\\", "\\\\", `"`, "\\\"")

	escapeQuotes := func(s string) string {
		return quoteEscaper.Replace(s)
	}

	h := make(textproto.MIMEHeader)
	h.Set("Content-Disposition",
		fmt.Sprintf(`form-data; name="%s"; filename="%s"`,
			escapeQuotes(fieldname), escapeQuotes(filename)))

	contentType, err := GetFileContentType(data)
	if err != nil {
		return nil, err
	}
	h.Set("Content-Type", contentType)
	return w.CreatePart(h)
}

func GetFileContentType(file io.Reader) (string, error) {

	// Only the first 512 bytes are used to sniff the content type.
	buffer := make([]byte, 512)

	_, err := file.Read(buffer)
	if err != nil {
		return "", err
	}

	// Use the net/http package's handy DectectContentType function. Always returns a valid
	// content-type by returning "application/octet-stream" if no others seemed to match.
	contentType := http.DetectContentType(buffer)

	return contentType, nil
}
