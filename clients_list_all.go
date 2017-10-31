package invoicexpress

import (
	"net/http"
	"net/url"
)

func (s *ClientsService) NewListAllRequest() ClientsListAllRequest {
	return ClientsListAllRequest{
		api:       s.api,
		method:    http.MethodGet,
		urlParams: ClientsListAllURLParams{},
		queryParams: ClientsListAllQueryParams{
			Page:    0,
			PerPage: 0,
		},
	}
}

type ClientsListAllRequest struct {
	api *API
	// queryParams ClientsListAllQueryParams
	// pathParams  ClientsListAllPathParams
	method      string
	urlParams   ClientsListAllURLParams
	queryParams ClientsListAllQueryParams
}

func (r *ClientsListAllRequest) Method() string {
	return r.method
}

func (r *ClientsListAllRequest) SetMethod(method string) {
	r.method = method
}

func (r *ClientsListAllRequest) URL() url.URL {
	path := "clients.json"
	return r.api.GetEndpointURL(path)
}

func (r *ClientsListAllRequest) Do() (ClientsListAllResponseBody, error) {
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

func (r *ClientsListAllRequest) NewResponseBody() *ClientsListAllResponseBody {
	return &ClientsListAllResponseBody{}
}

type ClientsListAllURLParams struct {
}

type ClientsListAllQueryParams struct {
	Page    int `schema:"page,omitempty"`
	PerPage int `schema:"per_page,omitempty"`
}

type ClientsListAllResponseBody struct {
	Clients    Clients    `json:"clients"`
	Pagination Pagination `json:"pagination"`
}
