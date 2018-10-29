package invoicexpress

import (
	"net/http"
	"net/url"
)

func (s *ClientsService) NewCreateRequest() ClientsCreateRequest {
	return ClientsCreateRequest{
		api:         s.api,
		method:      http.MethodPost,
		urlParams:   ClientsCreateURLParams{},
		requestBody: s.NewCreateRequestBody(),
	}
}

type ClientsCreateRequest struct {
	api *API
	// queryParams ClientsCreateQueryParams
	// pathParams  ClientsCreatePathParams
	method      string
	headers     http.Header
	urlParams   ClientsCreateURLParams
	requestBody ClientsCreateRequestBody
}

func (r *ClientsCreateRequest) Method() string {
	return r.method
}

func (r *ClientsCreateRequest) SetMethod(method string) {
	r.method = method
}

func (r *ClientsCreateRequest) URLParams() *ClientsCreateURLParams {
	return &r.urlParams
}

func (r *ClientsCreateRequest) RequestBody() *ClientsCreateRequestBody {
	return &r.requestBody
}

func (r *ClientsCreateRequest) SetRequestBody(body ClientsCreateRequestBody) {
	r.requestBody = body
}

func (r *ClientsCreateRequest) URL() url.URL {
	return r.api.GetEndpointURL("clients.json")
}

func (r *ClientsCreateRequest) Do() (ClientsCreateResponseBody, error) {
	req, err := r.api.NewRequest(nil, r.Method(), r.URL(), r.RequestBody())
	if err != nil {
		return *r.NewResponseBody(), err
	}

	responseBody := r.NewResponseBody()
	_, err = r.api.Do(req, responseBody)
	return *responseBody, err
}

func (r *ClientsCreateRequest) NewResponseBody() *ClientsCreateResponseBody {
	return &ClientsCreateResponseBody{}
}

func NewClientsCreateURLParams() *ClientsCreateURLParams {
	return &ClientsCreateURLParams{}
}

type ClientsCreateURLParams struct {
}

func (s *ClientsService) NewCreateRequestBody() ClientsCreateRequestBody {
	return ClientsCreateRequestBody{
		Client: NewClient{},
	}
}

type ClientsCreateRequestBody struct {
	Client NewClient `json:"client"`
}

type ClientsCreateResponseBody struct {
}
