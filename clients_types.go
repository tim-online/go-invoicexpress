package invoicexpress

type Clients []Client

type Client struct {
	ID               int         `json:"id"`
	Name             string      `json:"name"`
	Code             string      `json:"code"`
	Language         interface{} `json:"language"`
	Email            string      `json:"email"`
	Address          string      `json:"address"`
	City             string      `json:"city"`
	PostalCode       string      `json:"postal_code"`
	FiscalID         string      `json:"fiscal_id"`
	Website          string      `json:"website"`
	Country          string      `json:"country"`
	Phone            string      `json:"phone"`
	Fax              string      `json:"fax"`
	PreferredContact struct {
		Name   string      `json:"name"`
		Email  string      `json:"email"`
		Phone  string      `json:"phone"`
		Mobile interface{} `json:"mobile"`
	} `json:"preferred_contact"`
	Observations string      `json:"observations"`
	SendOptions  interface{} `json:"send_options"`
}

type NewClient struct {
	Name             string  `json:"name"`
	Code             string  `json:"code"`
	Email            string  `json:"email"`
	Address          string  `json:"address"`
	City             string  `json:"city"`
	PostalCode       string  `json:"postal_code"`
	Country          Country `json:"country"`
	FiscalID         string  `json:"fiscal_id"`
	Website          string  `json:"website"`
	Phone            string  `json:"phone"`
	Fax              string  `json:"fax"`
	Observations     string  `json:"observations"`
	PreferredContact Contact `json:"preferred_contact"`
}

type Contact struct {
	Name   string `json:"name"`
	Email  string `json:"email"`
	Phone  string `json:"phone"`
	Mobile string `json:"mobile"`
}
