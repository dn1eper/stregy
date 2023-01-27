package backtester

type BacktesterDTO struct {
	StrategyName        string `json:"strategy_name"`
	Timeframe           int    `json:"timeframe"`
	Symbol              string `json:"symbol"`
	StartDate           string `json:"start_date"`
	EndDate             string `json:"end_date"`
	HighOrderResolution bool   `json:"high_order_resolution"`
	BarsNeeded          int    `json:"bars_needed"`
	ATRperiod           int    `json:"atr_period,omitempty"`
}
