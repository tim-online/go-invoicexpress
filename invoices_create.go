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
	return InvoicesCreateRequestBody{}
}

type InvoicesCreateRequestBody struct {
	Invoice           *NewInvoice `json:"invoice"`
	SimplifiedInvoice *NewInvoice `json:"simplified_invoice"`
	InvoiceReceipt    *NewInvoice `json:"invoice_receipt"`
	CreditNote        *NewInvoice `json:"credit_note"`
	DebitNote         *NewInvoice `json:"debit_note"`
}

type InvoicesCreateResponseBody struct {
	Invoice    ResponseInvoice `json:"invoice"`
	CreditNote ResponseInvoice `json:"credit_note"`
}

type ResponseInvoice struct {
	ID                     int            `json:"id"`
	Status                 string         `json:"status"`
	Archived               bool           `json:"archived"`
	Type                   string         `json:"type"`
	SequenceNumber         string         `json:"sequence_number"`
	InvertedSequenceNumber string         `json:"inverted_sequence_number"`
	SequenceID             string         `json:"sequence_id"`
	Date                   string         `json:"date"`
	DueDate                string         `json:"due_date"`
	Reference              string         `json:"reference"`
	Observations           string         `json:"observations"`
	Retention              interface{}    `json:"retention"`
	Permalink              string         `json:"permalink"`
	Sum                    float64        `json:"sum"`
	Discount               float64        `json:"discount"`
	BeforeTaxes            float64        `json:"before_taxes"`
	Taxes                  float64        `json:"taxes"`
	Total                  float64        `json:"total"`
	Currency               string         `json:"currency"`
	Client                 ResponseClient `json:"client"`
	Items                  []ResponseItem `json:"items"`
}

type ResponseClient struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
	Country string `json:"country"`
	Code    string `json:"code"`
	City    string `json:"city"`
}

type ResponseItem struct {
	Name           string  `json:"name"`
	Description    string  `json:"description"`
	UnitPrice      Number  `json:"unit_price"`
	Unit           string  `json:"unit"`
	Quantity       Number  `json:"quantity"`
	Tax            Tax     `json:"tax"`
	Discount       float64 `json:"discount"`
	Subtotal       float64 `json:"subtotal"`
	TaxAmount      float64 `json:"tax_amount"`
	DiscountAmount float64 `json:"discount_amount"`
	Total          float64 `json:"total"`
}
