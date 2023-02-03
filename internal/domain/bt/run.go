package bt

import (
	"fmt"
	"os"
	"path"
	"stregy/internal/domain/order"
	"stregy/internal/domain/quote"
	strategy1 "stregy/internal/domain/strategy"
	"stregy/pkg/logging"
)

var strategy strategy1.Strategy

func (backtest *Backtest) RunOnQuotes(s strategy1.Strategy, quotes <-chan quote.Quote, firstQuote quote.Quote) error {
	loggingConfig = LoggingConfig{LogOrderStatusChange: false, PricePrecision: backtest.Symbol.Precision}
	InitializeBacktester(s)

	Time = firstQuote.Time
	Price = firstQuote.Open
	InitLogger(backtest.ID + ".log")
	Printf("running backtest with strategy strat1 on period period [%s; %s]", backtest.StartTime.Format("2006-01-02 15:04:05"), backtest.EndTime.Format("2006-01-02 15:04:05"))

	quoteGen, err := NewQuoteGenerator(s, backtest.TimeframeSec, firstQuote)
	if err != nil {
		return err
	}

	qCount := 0
	for q := range quotes {
		Time = q.Time
		Price = q.Close

		for _, o := range orders {
			if o.Type == order.Limit {
				if o.Diraction == order.Long {
					if q.Low <= o.Price {
						executeOrder(o, o.Price)
						continue
					}
				} else {
					if q.High >= o.Price {
						executeOrder(o, o.Price)
						continue
					}
				}
			} else if o.Type == order.StopMarket {
				if o.Price >= q.Low && o.Price <= q.High {
					executeOrder(o, q.Close)
					continue
				}
			} else if o.Type == order.Market {
				executeOrder(o, q.Close)
				continue
			}
		}

		// debug breaker
		qCount += 1
		if qCount == 3601 {
			break
		}

		quoteGen.OnQuote(q)
	}

	wd, _ := os.Getwd()
	reportDir := path.Join(wd, "reports")
	os.Mkdir(reportDir, os.ModePerm)
	err = backtest.AccountHistoryService.CreateReport(orderHistory, backtest.Symbol, path.Join(reportDir, backtest.ID+".csv"))
	if err != nil {
		logging.GetLogger().Error(fmt.Sprintf("Error creating backtest report: %v", err))
	}

	return nil
}

func InitializeBacktester(s strategy1.Strategy) {
	strategy = s
	orders = make(map[int64]*order.Order)
	positions = make(map[int64]*order.Position)
}

func Terminate() {
	Print("Terminated")
	os.Exit(0)
}
