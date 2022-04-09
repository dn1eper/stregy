package exgaccount

type CreateExchangeAccountDTO struct {
	ExchangeID       string `json:"exchange_id,omitempty"`
	ConnectionString string `json:"connection_string"`
	Name             string `json:"name"`
}
