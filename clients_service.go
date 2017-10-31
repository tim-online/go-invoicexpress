package invoicexpress

func NewClientsService(api *API) *ClientsService {
	return &ClientsService{api: api}
}

type ClientsService struct {
	api *API
}
