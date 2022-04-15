package stratexec

type CreateStrategyExecutionDTO struct {
	StrategyID        string `json:"strategy_id"`
	ExchangeAccountID string `json:"exchange_account_id"`
	Timeframe         string `json:"timeframe"`
	Symbol            string `json:"symbol"`
	StartTime         string `json:"start_time"`
	EndTime           string `json:"end_time"`
}
