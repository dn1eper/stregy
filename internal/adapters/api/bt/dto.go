package bt

type BacktesterDTO struct {
	StrategyName string `json:"strategy_name"`
	Symbol       string `json:"symbol"`
	TimeframeSec int    `json:"timeframe_sec"`
	StartDate    string `json:"start_date"`
	EndDate      string `json:"end_date"`
}
