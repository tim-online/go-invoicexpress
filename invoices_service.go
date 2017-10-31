package invoicexpress

func NewInvoicesService(api *API) *InvoicesService {
	return &InvoicesService{api: api}
}

type InvoicesService struct {
	api *API
}
