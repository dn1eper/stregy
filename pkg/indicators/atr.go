package indicators

import (
	"stregy/internal/domain/quote"
	"stregy/pkg/utils"
)

func _ATR(period int, quotes []quote.Quote) float64 {
	if len(quotes) < period {
		return 0
	}

	sum := 0.0
	for i := len(quotes) - 1; i >= 0; i-- {
		sum += TR(quotes, i)
	}

	return sum / float64(period)
}

func _ATRinc(prevATR float64, period int, quotes []quote.Quote) float64 {
	return (prevATR*float64(period-1) + TR(quotes, len(quotes)-1)) / float64(period)
}

func TR(quotes []quote.Quote, idx int) float64 {
	if idx-1 < 0 {
		return 0
	}

	return utils.Max(
		quotes[idx].High-quotes[idx].Low,
		quotes[idx].High-quotes[idx-1].Close,
		quotes[idx-1].Close-quotes[idx].Low,
	)
}
