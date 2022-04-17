package backtester

type BacktestDTO struct {
	StrategyID        string `json:"strategy_id"`
	ExchangeAccountID string `json:"exchange_account_id"`
	Timeframe         string `json:"timeframe"`
	Symbol            string `json:"symbol"`
	StartDate         string `json:"start_time"`
	EndDate           string `json:"end_time"`
}
