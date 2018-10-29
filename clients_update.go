package invoicexpress

import (
	"net/http"
	"net/url"
	"strings"
)

func (s *ClientsService) NewUpdateRequest() ClientsUpdateRequest {
	return ClientsUpdateRequest{
		api:         s.api,
		method:      http.MethodPut,
		urlParams:   ClientsUpdateURLParams{},
		requestBody: s.NewUpdateRequestBody(),
	}
}

type ClientsUpdateRequest struct {
	api         *API
	method      string
	headers     http.Header
	urlParams   ClientsUpdateURLParams
	requestBody ClientsUpdateRequestBody
}

func (r *ClientsUpdateRequest) Method() string {
	return r.method
}

func (r *ClientsUpdateRequest) SetMethod(method string) {
	r.method = method
}

func (r *ClientsUpdateRequest) URLParams() *ClientsUpdateURLParams {
	return &r.urlParams
}

func (r *ClientsUpdateRequest) RequestBody() *ClientsUpdateRequestBody {
	return &r.requestBody
}

func (r *ClientsUpdateRequest) SetRequestBody(body ClientsUpdateRequestBody) {
	r.requestBody = body
}

func (r *ClientsUpdateRequest) URL() url.URL {
	path := "clients/:client-id.json"
	path = strings.Replace(path, ":client-id", r.urlParams.ClientID, 1)
	return r.api.GetEndpointURL(path)
}

func (r *ClientsUpdateRequest) Do() (ClientsUpdateResponseBody, error) {
	req, err := r.api.NewRequest(nil, r.Method(), r.URL(), r.RequestBody())
	if err != nil {
		return *r.NewResponseBody(), err
	}

	responseBody := r.NewResponseBody()
	_, err = r.api.Do(req, responseBody)
	return *responseBody, err
}

func (r *ClientsUpdateRequest) NewResponseBody() *ClientsUpdateResponseBody {
	return &ClientsUpdateResponseBody{}
}

func NewClientsUpdateURLParams() *ClientsUpdateURLParams {
	return &ClientsUpdateURLParams{}
}

type ClientsUpdateURLParams struct {
	ClientID string
}

func (s *ClientsService) NewUpdateRequestBody() ClientsUpdateRequestBody {
	return ClientsUpdateRequestBody{
		Client: NewClient{},
	}
}

type ClientsUpdateRequestBody struct {
	Client NewClient `json:"client"`
}

type ClientsUpdateResponseBody struct {
}
