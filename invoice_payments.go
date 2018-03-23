package invoicexpress

import (
	"fmt"
	"net/http"
	"net/url"
)

func (s *InvoicesService) NewPaymentRequest() InvoicePaymentRequest {
	return InvoicePaymentRequest{
		api:         s.api,
		method:      http.MethodPost,
		params:      *NewInvoicePaymentParams(),
		requestBody: s.NewPaymentRequestBody(),
	}
}

type InvoicePaymentRequest struct {
	api         *API
	method      string
	headers     http.Header
	params      InvoicePaymentParams
	requestBody InvoicePaymentRequestBody
}

func (r *InvoicePaymentRequest) Method() string {
	return r.method
}

func (r *InvoicePaymentRequest) SetMethod(method string) {
	r.method = method
}

func (r *InvoicePaymentRequest) Params() *InvoicePaymentParams {
	return &r.params
}

func (r *InvoicePaymentRequest) RequestBody() *InvoicePaymentRequestBody {
	return &r.requestBody
}

func (r *InvoicePaymentRequest) SetRequestBody(body InvoicePaymentRequestBody) {
	r.requestBody = body
}

func (r *InvoicePaymentRequest) URL() url.URL {
	path := fmt.Sprintf("documents/%v/partial_payments.json", r.params.invoiceID)
	return r.api.GetEndpointURL(path)
}

func (r *InvoicePaymentRequest) Do() (InvoicePaymentResponseBody, error) {
	req, err := r.api.NewRequest(nil, r.Method(), r.URL(), r.RequestBody())
	if err != nil {
		return *r.NewResponseBody(), err
	}

	responseBody := r.NewResponseBody()
	_, err = r.api.Do(req, responseBody)
	return *responseBody, err
}

func (r *InvoicePaymentRequest) NewResponseBody() *InvoicePaymentResponseBody {
	return &InvoicePaymentResponseBody{}
}

func NewInvoicePaymentParams() *InvoicePaymentParams {
	return &InvoicePaymentParams{
		documentType: Invoice,
	}
}

type InvoicePaymentParams struct {
	documentType InvoiceDocumentType
	invoiceID    int
}

func (p *InvoicePaymentParams) SetInvoiceID(invoiceID int) {
	p.invoiceID = invoiceID
}

func (s *InvoicesService) NewPaymentRequestBody() InvoicePaymentRequestBody {
	return InvoicePaymentRequestBody{
		PartialPayment: PaymentItem{},
	}
}

type InvoicePaymentRequestBody struct {
	PartialPayment PaymentItem `json:"partial_payment"`
}

type PaymentItem struct {
	PaymentMechanism string  `json:"payment_mechanism"`
	Amount           float64 `json:"amount"`
	PaymentDate      string  `json:"payment_date"`
}

type InvoicePaymentResponseBody struct {
}
