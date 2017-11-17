package invoicexpress

import (
	"fmt"
	"net/http"
	"net/url"
)

func (s *InvoicesService) NewCreateRequest() InvoicesCreateRequest {
	return InvoicesCreateRequest{
		api:         s.api,
		method:      http.MethodPost,
		params:      *NewInvoicesCreateParams(),
		requestBody: s.NewCreateRequestBody(),
	}
}

type InvoicesCreateRequest struct {
	api *API
	// queryParams InvoicesCreateQueryParams
	// pathParams  InvoicesCreatePathParams
	method      string
	headers     http.Header
	params      InvoicesCreateParams
	requestBody InvoicesCreateRequestBody
}

func (r *InvoicesCreateRequest) Method() string {
	return r.method
}

func (r *InvoicesCreateRequest) SetMethod(method string) {
	r.method = method
}

func (r *InvoicesCreateRequest) Params() *InvoicesCreateParams {
	return &r.params
}

func (r *InvoicesCreateRequest) RequestBody() *InvoicesCreateRequestBody {
	return &r.requestBody
}

func (r *InvoicesCreateRequest) SetRequestBody(body InvoicesCreateRequestBody) {
	r.requestBody = body
}

func (r *InvoicesCreateRequest) URL() url.URL {
	path := fmt.Sprintf("%s.json", r.params.documentType.Path)
	return r.api.GetEndpointURL(path)
}

func (r *InvoicesCreateRequest) Do() (InvoicesCreateResponseBody, error) {
	req, err := r.api.NewRequest(nil, r.Method(), r.URL(), r.RequestBody())
	if err != nil {
		return *r.NewResponseBody(), err
	}

	responseBody := r.NewResponseBody()
	_, err = r.api.Do(req, responseBody)
	return *responseBody, err
}

func (r *InvoicesCreateRequest) NewResponseBody() *InvoicesCreateResponseBody {
	return &InvoicesCreateResponseBody{}
}

func NewInvoicesCreateParams() *InvoicesCreateParams {
	return &InvoicesCreateParams{
		documentType: Invoice,
	}
}

type InvoicesCreateParams struct {
	documentType InvoiceDocumentType
}

func (p *InvoicesCreateParams) SetDocumentType(documentType InvoiceDocumentType) {
	p.documentType = documentType
}

func (s *InvoicesService) NewCreateRequestBody() InvoicesCreateRequestBody {
	return InvoicesCreateRequestBody{
		Invoice: NewInvoice{
			Items: []NewItem{},
		},
	}
}

type InvoicesCreateRequestBody struct {
	Invoice NewInvoice `json:"invoice"`
}

type InvoicesCreateResponseBody struct {
}
