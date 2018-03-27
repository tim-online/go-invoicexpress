package invoicexpress

import "github.com/aodin/date"

type InvoiceDocumentType struct {
	Type string
	Path string
}

var (
	Invoice           = InvoiceDocumentType{"Invoice", "invoices"}
	SimplifiedInvoice = InvoiceDocumentType{"SimplifiedInvoice", "simplified_invoices"}
	InvoiceReceipt    = InvoiceDocumentType{"InvoiceReceipt", "invoice_receipts"}
	CreditNote        = InvoiceDocumentType{"CreditNote", "credit_notes"}
	DebitNote         = InvoiceDocumentType{"DebitNote", "debit_notes"}
)

func (d InvoiceDocumentType) String() string {
	return d.Type
}

type NewInvoice struct {
	Date                 string           `json:"date"`
	DueDate              strubg           `json:"due_date"`
	Reference            string           `json:"reference"`
	Observations         string           `json:"observations"`
	Retention            string           `json:"retention"`
	TaxExemption         TaxExemptionCode `json:"tax_exemption,omitempty"`
	SequenceID           string           `json:"sequence_id"`
	ManualSequenceNumber string           `json:"manual_sequence_number"`
	Client               NewClient        `json:"client"`
	Items                []NewItem        `json:"items"`
	PartialPayments      []PaymentItem    `json:"partial_payments"`
	MbReference          string           `json:"mb_reference"`
	OwnerInvoiceID       string           `json:"owner_invoice_id"`
}

type NewItem struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	UnitPrice   float64 `json:"unit_price"`
	Quantity    float64 `json:"quantity"`
	Unit        string  `json:"unit"`
	Discount    float64 `json:"discount"`
	Tax         Tax     `json:"tax"`
}

type Tax struct {
	Name string `json:"name"`
}

type Country string

type InvoiceStatus string

var (
	Draft      InvoiceStatus = "draft"
	Sent       InvoiceStatus = "sent"
	Settled    InvoiceStatus = "settled"
	Finalized  InvoiceStatus = "finalized"
	Canceled   InvoiceStatus = "canceled"
	SecondCopy InvoiceStatus = "second_copy"
)
