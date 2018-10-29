package invoicexpress

type Clients []Client

type Client struct {
	ID int `json:"id"`
	NewClient
}

type NewClient struct {
	// [Required] Client name, normally used for a company name.
	Name string `json:"name"`

	// [Required] Client code, your specific code for the client.
	Code string `json:"code"`

	// [Optional] Client email address. Must be a valid email address ex:
	// foo@bar.com
	Email string `json:"email,omitempty"`

	// [Optional] Client language. May be 'en', 'pt' or' es'. Defaults to the
	// account language.
	Language string `json:"language,omitempty"`

	// [Optional] Client company address.
	Address string `json:"address,omitempty"`

	// [Optional] Client’s city.
	City string `json:"city"`

	// [Optional] Client’s postal code for it’s company address.
	PostalCode string `json:"postal_code,omitempty"`

	// [Optional] Country, normally used for a company country. Although country
	// is optional, when supplied, it should match one of the country list on
	// the Appendix of this documentation.
	Country Country `json:"country"`

	// [Optional] The client fiscal ID (Número de Contribuinte)
	FiscalID string `json:"fiscal_id,omitempty"`

	// [Optional] The client website address.
	Website string `json:"website,omitempty"`

	// [Optional] The client phone number.
	Phone string `json:"phone,omitempty"`

	// [Optional] The client fax number
	Fax string `json:"fax"`

	// [Optional] Preferred contact details
	PreferredContact Contact `json:"preferred_contact,omitempty"`

	// [Optional] Default observations for the client, the text in here will be
	// added to the observations field of all the invoices sent to this client.
	Observations string `json:"observations,omitempty"`

	// [Optional] Send options for the client. Available options are: 1 - send
	// only the original document. 2 - the original and duplicate. 3 - the
	// original, duplicate and triplicate. These affect the generated pdf.
	SendOptions string `json:"send_options,omitempty"`
}

type Contact struct {
	Name   string `json:"name"`
	Email  string `json:"email"`
	Phone  string `json:"phone"`
	Mobile string `json:"mobile"`
}
