package invoicexpress_test

import (
	"log"
	"os"
	"testing"

	invoicexpress "github.com/tim-online/go-invoicexpress"
)

func TestClientsListAll(t *testing.T) {
	accountName := os.Getenv("INVOICEXPRESS_ACCOUNTNAME")
	token := os.Getenv("INVOICEXPRESS_TOKEN")
	api := invoicexpress.NewAPI(nil, accountName, token)
	api.SetDebug(true)

	req := api.Clients.NewListAllRequest()
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}
	log.Println(resp)
}
