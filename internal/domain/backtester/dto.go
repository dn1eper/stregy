package backtester

type BacktesterDTO struct {
	StrategyID        string `json:"strategy_id"`
	ExchangeAccountID string `json:"exchange_account_id"`
	Timeframe         int    `json:"timeframe"`
	Symbol            string `json:"symbol"`
	StartDate         string `json:"start_date"`
	EndDate           string `json:"end_date"`
}
