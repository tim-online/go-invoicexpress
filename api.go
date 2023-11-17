package invoicexpress

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

const (
	libraryVersion = "0.0.1"
	userAgent      = "go-invoicexpress/" + libraryVersion
	mediaType      = "application/json"
	charset        = "utf-8"
)

var (
	BaseURL = url.URL{
		Scheme: "https",
		Host:   "{account_name}.app.invoicexpress.com",
		Path:   "",
	}
)

// NewAPI returns a new InvoiceXpress API client
func NewAPI(httpClient *http.Client, accountName string, token string) *API {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	api := &API{
		http: httpClient,
	}

	api.SetAccountName(accountName)
	api.SetToken(token)
	api.SetBaseURL(BaseURL)
	api.SetDebug(false)
	api.SetUserAgent(userAgent)
	api.SetMediaType(mediaType)
	api.SetCharset(charset)

	// Services
	api.Clients = NewClientsService(api)
	api.Invoices = NewInvoicesService(api)

	return api
}

// API manages communication with InvoiceXpress API
type API struct {
	// HTTP client used to communicate with the API.
	http *http.Client

	debug   bool
	baseURL url.URL

	// credentials
	accountName string
	token       string

	// User agent for client
	userAgent string

	mediaType string
	charset   string

	// Optional function called after every successful request made to the DO APIs
	onRequestCompleted RequestCompletionCallback

	// Services used for communicating with the API
	Clients  *ClientsService
	Invoices *InvoicesService
}

// RequestCompletionCallback defines the type of the request callback function
type RequestCompletionCallback func(*http.Request, *http.Response)

func (api *API) Debug() bool {
	return api.debug
}

func (api *API) SetDebug(debug bool) {
	api.debug = debug
}

func (api *API) AccountName() string {
	return api.accountName
}

func (api *API) SetAccountName(accountName string) {
	api.accountName = accountName
}

func (api *API) Token() string {
	return api.token
}

func (api *API) SetToken(token string) {
	api.token = token
}

func (api *API) BaseURL() url.URL {
	return api.baseURL
}

func (api *API) SetBaseURL(baseURL url.URL) {
	api.baseURL = baseURL
}

func (api *API) SetMediaType(mediaType string) {
	api.mediaType = mediaType
}

func (api *API) MediaType() string {
	return mediaType
}

func (api *API) SetCharset(charset string) {
	api.charset = charset
}

func (api *API) Charset() string {
	return charset
}

func (api *API) SetUserAgent(userAgent string) {
	api.userAgent = userAgent
}

func (api *API) UserAgent() string {
	return userAgent
}

func (api *API) GetEndpointURL(path string) url.URL {
	apiURL := api.BaseURL()
	apiURL.Host = strings.Replace(apiURL.Host, "{account_name}", api.AccountName(), 1)
	apiURL.Path = apiURL.Path + path
	return apiURL
}

func (api *API) NewRequest(ctx context.Context, method string, URL url.URL, body interface{}) (*http.Request, error) {
	// convert body struct to json
	buf := new(bytes.Buffer)
	if body != nil {
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	// add api key to url
	q := URL.Query()
	q.Add("api_key", api.Token())
	URL.RawQuery = q.Encode()

	// create new http request
	req, err := http.NewRequest(method, URL.String(), buf)
	if err != nil {
		return nil, err
	}

	// req.SetBasicAuth(c.Username(), c.Password())

	// optionally pass along context
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	// set other headers
	req.Header.Add("Content-Type", fmt.Sprintf("%s; charset=%s", api.MediaType(), api.Charset()))
	req.Header.Add("Accept", api.MediaType())
	req.Header.Add("User-Agent", api.UserAgent())

	return req, nil
}

// Do sends an API request and returns the API response. The API response is json decoded and stored in the value
// pointed to by v, or returned as an error if an API error has occurred. If v implements the io.Writer interface,
// the raw response will be written to v, without attempting to decode it.
func (api *API) Do(req *http.Request, responseBody interface{}) (*http.Response, error) {
	if api.debug == true {
		dump, _ := httputil.DumpRequestOut(req, true)
		log.Println(string(dump))
	}

	httpResp, err := api.http.Do(req)
	if err != nil {
		return nil, err
	}

	if api.onRequestCompleted != nil {
		api.onRequestCompleted(req, httpResp)
	}

	// close body io.Reader
	defer func() {
		if rerr := httpResp.Body.Close(); err == nil {
			err = rerr
		}
	}()

	if api.debug == true {
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

	// interface implements io.Writer: write Body to it
	// if w, ok := response.Envelope.(io.Writer); ok {
	// 	_, err := io.Copy(w, httpResp.Body)
	// 	return httpResp, err
	// }

	// try to decode body into interface parameter
	if responseBody != nil {
		err = json.NewDecoder(httpResp.Body).Decode(responseBody)
		if err != nil && err != io.EOF {
			// create a simple error response
			errorResponse := &ErrorResponse{Response: httpResp}
			errorResponse.Errors = append(errorResponse.Errors, err)
			return httpResp, errorResponse
		}
	}

	return httpResp, nil
}

// CheckResponse checks the API response for errors, and returns them if
// present. A response is considered an error if it has a status code outside
// the 200 range. API error responses are expected to have either no response
// body, or a XML response body that maps to ErrorResponse. Any other response
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

	err := checkContentType(r)
	if err != nil {
		errorResponse.Errors = append(errorResponse.Errors, errors.New(r.Status))
		return errorResponse
	}

	// read data and copy it back
	data, err := ioutil.ReadAll(r.Body)
	r.Body = ioutil.NopCloser(bytes.NewReader(data))
	if err != nil {
		return errorResponse
	}

	if len(data) == 0 {
		return errorResponse
	}

	// convert json to struct
	err = json.Unmarshal(data, errorResponse)
	if err != nil {
		errorResponse.Errors = append(errorResponse.Errors, err)
		return errorResponse
	}

	return errorResponse
}

type ErrorResponse struct {
	// HTTP response that caused this error
	Response *http.Response `json:"-"`

	Errors []error
}

// {
//   "errors": [
//     {
//       "error": "Document should have at least one item"
//     },
//     {
//       "error": "Date is not valid"
//     },
//     {
//       "error": "Tax exemption should be one of: M09, M08, M07, M06, M05, M04, M03, M02, M01, M16, M15, M14, M13, M12, M11, M10, M99"
//     }
//   ]
// }

func (r *ErrorResponse) UnmarshalJSON(data []byte) error {
	tmp := struct {
		Errors []struct {
			Error string `json:"error"`
		} `json:"errors"`
	}{}

	err := json.Unmarshal(data, &tmp)
	if err == nil {
		for _, err := range tmp.Errors {
			r.Errors = append(r.Errors, errors.New(err.Error))
		}
		return nil
	}

	tmp2 := [][]string{}

	err = json.Unmarshal(data, &tmp2)
	if err == nil {
		for _, v := range tmp2 {
			for _, e := range v {
				r.Errors = append(r.Errors, errors.New(e))
			}
		}
		return nil
	}

	// {"errors":[{"error":["CantHandle: draft for CreditNote"]}]}
	tmp3 := struct {
		Errors []struct {
			Error []string `json:"error"`
		} `json:"errors"`
	}{}

	err = json.Unmarshal(data, &tmp3)
	if err == nil {
		for _, v := range tmp3.Errors {
			for _, e := range v.Error {
				r.Errors = append(r.Errors, errors.New(e))
			}
		}
		return nil
	}

	// {"errors":{"error":"Document creation limit reached for the period from 15/12/2022 to 15/01/2023."}}
	tmp4 := struct {
		Errors struct {
			Error string `json:"error"`
		} `json:"errors"`
	}{}

	err = json.Unmarshal(data, &tmp4)
	if err == nil {
		r.Errors = append(r.Errors, errors.New(tmp4.Errors.Error))
		return nil
	}

	// {"errors":["Fiscal is invalid"]}
	tmp5 := struct {
		Errors []string `json:"errors"`
	}{}

	err = json.Unmarshal(data, &tmp5)
	if err == nil {
		for _, e := range tmp5.Errors {
			r.Errors = append(r.Errors, errors.New(e))
		}
		return nil
	}

	return err
}

func (r ErrorResponse) Error() string {
	if len(r.Errors) > 0 {
		str := []string{}
		for _, err := range r.Errors {
			str = append(str, err.Error())
		}
		return strings.Join(str, ", ")
	}

	switch r.Response.StatusCode {
	case 401:
		return "The API Key parameter is missing or is incorrectly entered."
	case 404:
		return "The requested resource does not exist."
	case 406:
		return "The :document-id provided is in an invalid state."
	case 422:
		return "Some parameters were incorrect."
	}

	return fmt.Sprintf("Unknown status code %d", r.Response.StatusCode)
}

func checkContentType(response *http.Response) error {
	// check content-type (application/soap+xml; charset=utf-8)
	header := response.Header.Get("Content-Type")
	contentType := strings.Split(header, ";")[0]
	if contentType != mediaType {
		return fmt.Errorf("Expected Content-Type \"%s\", got \"%s\"", mediaType, contentType)
	}

	return nil
}
