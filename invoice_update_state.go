package invoicexpress

import (
	"fmt"
	"net/http"
	"net/url"
)

func (s *InvoicesService) NewUpdateStateRequest() InvoiceUpdateStateRequest {
	return InvoiceUpdateStateRequest{
		api:         s.api,
		method:      http.MethodPost,
		params:      *NewInvoiceUpdateStateParams(),
		requestBody: s.NewUpdateStateRequestBody(),
	}
}

type InvoiceUpdateStateRequest struct {
	api         *API
	method      string
	headers     http.Header
	params      InvoiceUpdateStateParams
	requestBody InvoiceUpdateStateRequestBody
}

func (r *InvoiceUpdateStateRequest) Method() string {
	return r.method
}

func (r *InvoiceUpdateStateRequest) SetMethod(method string) {
	r.method = method
}

func (r *InvoiceUpdateStateRequest) Params() *InvoiceUpdateStateParams {
	return &r.params
}

func (r *InvoiceUpdateStateRequest) RequestBody() *InvoiceUpdateStateRequestBody {
	return &r.requestBody
}

func (r *InvoiceUpdateStateRequest) SetRequestBody(body InvoiceUpdateStateRequestBody) {
	r.requestBody = body
}

func (r *InvoiceUpdateStateRequest) SetState(state InvoiceStatus) {
	r.requestBody.Invoice.State = state
}

func (r *InvoiceUpdateStateRequest) URL() url.URL {
	path := fmt.Sprintf("%[1]s/%[2]v/change-state.json", r.params.documentType.Path, r.params.invoiceID)
	return r.api.GetEndpointURL(path)
}

func (r *InvoiceUpdateStateRequest) Do() (InvoiceUpdateStateResponseBody, error) {
	req, err := r.api.NewRequest(nil, r.Method(), r.URL(), r.RequestBody())
	if err != nil {
		return *r.NewResponseBody(), err
	}

	responseBody := r.NewResponseBody()
	_, err = r.api.Do(req, responseBody)
	return *responseBody, err
}

func (r *InvoiceUpdateStateRequest) NewResponseBody() *InvoiceUpdateStateResponseBody {
	return &InvoiceUpdateStateResponseBody{}
}

func NewInvoiceUpdateStateParams() *InvoiceUpdateStateParams {
	return &InvoiceUpdateStateParams{
		documentType: Invoice,
	}
}

type InvoiceUpdateStateParams struct {
	documentType InvoiceDocumentType
	invoiceID    int
}

func (p *InvoiceUpdateStateParams) SetDocumentType(documentType InvoiceDocumentType) {
	p.documentType = documentType
}

func (p *InvoiceUpdateStateParams) SetInvoiceID(invoiceID int) {
	p.invoiceID = invoiceID
}

func (s *InvoicesService) NewUpdateStateRequestBody() InvoiceUpdateStateRequestBody {
	return InvoiceUpdateStateRequestBody{
		Invoice: UpdateStateBody{},
	}
}

type InvoiceUpdateStateRequestBody struct {
	Invoice UpdateStateBody `json:"invoice"`
}

type UpdateStateBody struct {
	State InvoiceStatus `json:"state"`
}

type InvoiceUpdateStateResponseBody struct {
}
