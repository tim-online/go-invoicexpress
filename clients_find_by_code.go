package invoicexpress

import (
	"net/http"
	"net/url"
)

func (s *ClientsService) NewFindByCodeRequest() ClientsFindByCodeRequest {
	return ClientsFindByCodeRequest{
		api:         s.api,
		method:      http.MethodGet,
		urlParams:   ClientsFindByCodeURLParams{},
		queryParams: ClientsFindByCodeQueryParams{},
	}
}

type ClientsFindByCodeRequest struct {
	api *API
	// queryParams ClientsFindByCodeQueryParams
	// pathParams  ClientsFindByCodePathParams
	method      string
	urlParams   ClientsFindByCodeURLParams
	queryParams ClientsFindByCodeQueryParams
}

func (r *ClientsFindByCodeRequest) Method() string {
	return r.method
}

func (r *ClientsFindByCodeRequest) SetMethod(method string) {
	r.method = method
}

func (r *ClientsFindByCodeRequest) URL() url.URL {
	path := "clients/find-by-code.json"
	return r.api.GetEndpointURL(path)
}

func (r *ClientsFindByCodeRequest) Do() (ClientsFindByCodeResponseBody, error) {
	// Create http request
	req, err := r.api.NewRequest(nil, r.Method(), r.URL(), nil)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	// Process query parameters
	err = AddQueryParamsToRequest(r.queryParams, req, true)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	responseBody := r.NewResponseBody()
	_, err = r.api.Do(req, responseBody)
	return *responseBody, err
}

func (r *ClientsFindByCodeRequest) NewResponseBody() *ClientsFindByCodeResponseBody {
	return &ClientsFindByCodeResponseBody{}
}

func (r *ClientsFindByCodeRequest) QueryParams() *ClientsFindByCodeQueryParams {
	return &r.queryParams
}

type ClientsFindByCodeURLParams struct {
}

type ClientsFindByCodeQueryParams struct {
	ClientCode string `schema:"client_code"`
}

type ClientsFindByCodeResponseBody struct {
	Client Client `json:"client"`
}
