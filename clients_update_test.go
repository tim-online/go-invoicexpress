package invoicexpress_test

import (
	"log"
	"os"
	"testing"

	invoicexpress "github.com/tim-online/go-invoicexpress"
)

func TestClientsUpdate(t *testing.T) {
	accountName := os.Getenv("INVOICEXPRESS_ACCOUNTNAME")
	token := os.Getenv("INVOICEXPRESS_TOKEN")
	api := invoicexpress.NewAPI(nil, accountName, token)
	api.SetDebug(true)

	req := api.Clients.NewUpdateRequest()
	req.URLParams().ClientID = "4847252"
	req.SetRequestBody(invoicexpress.ClientsUpdateRequestBody{
		Client: invoicexpress.NewClient{
			Name:       "Omniboost BV",
			Code:       "1234",
			Email:      "info@omniboost.io",
			Country:    "Netherlands",
			City:       "Zaamslag",
			PostalCode: "4543CJ",
			Website:    "https://omniboost.io",
			Phone:      "",
			Address:    "Axelsestraat 4",
			Language:   "EN",
			PreferredContact: invoicexpress.Contact{
				Name:  "Leon Bogaert",
				Email: "leon@omniboost.io",
			},
		},
	})
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}
	log.Println(resp)
}
