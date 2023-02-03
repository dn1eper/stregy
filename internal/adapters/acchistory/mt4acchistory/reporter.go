package mt4acchistory

import (
	"fmt"
	"os"
	"strconv"
	"stregy/internal/domain/bt"
	"stregy/internal/domain/order"
	"stregy/internal/domain/symbol"
)

type accountHistoryReporter struct {
}

func NewAccountHistoryReporter() bt.AccountHistoryReport {
	return accountHistoryReporter{}
}

// not correct: prints multiple rows per position
func (reporter accountHistoryReporter) CreateReport(orders []*order.Order, s symbol.Symbol, filePath string) error {
	f, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer f.Close()

	for _, o := range orders {
		positionOpenPrice := o.ExecutionPrice
		positionClosePrice := getPositionClosePrice(o.Position)
		var profit float64
		if o.Diraction == order.Long {
			profit = o.Size * (positionClosePrice - positionOpenPrice)
		} else {
			profit = o.Size * (positionClosePrice - positionOpenPrice)
		}

		s := fmt.Sprintf(
			"%d,%s,%s,%f,%s,%s,%s,%s,%s,%s,",
			o.ID,
			o.SubmissionTime.Format("2006-01-02 15:04:05"),
			formatOrderType(o),
			o.Size,
			s.Name,
			FormatPrice(positionOpenPrice, s.Precision),
			FormatPrice(getPositionStopLossPrice(o.Position), s.Precision),
			FormatPrice(getPositionTakeProfitPrice(o.Position), s.Precision),
			o.FCTime.Format("2006-01-02 15:04:05"),
			FormatPrice(positionClosePrice, s.Precision),
		)

		if o.Status == order.Filled {
			s += fmt.Sprintf("%d,%d,%d,%f\n", 0, 0, 0, profit)
		} else {
			s += "cancelled,,,\n"
		}

		f.WriteString(s)
	}

	return nil
}

func formatOrderType(o *order.Order) string {
	var res string
	if o.Diraction == order.Long {
		res += "buy"
	} else {
		res += "sell"
	}
	if o.Type == order.Limit {
		res += " limit"
	}

	return res
}

func getPositionStopLossPrice(p *order.Position) float64 {
	var res float64
	if p.StopOrder != nil {
		res = p.StopOrder.Price
	}

	return res
}

func getPositionTakeProfitPrice(p *order.Position) float64 {
	var res float64
	if p.TakeOrder != nil {
		res = p.TakeOrder.Price
	}

	return res
}

func getPositionClosePrice(p *order.Position) float64 {
	if p.TakeOrder != nil && p.TakeOrder.Status == order.Filled {
		return p.TakeOrder.ExecutionPrice
	}
	if p.StopOrder != nil && p.StopOrder.Status == order.Filled {
		return p.StopOrder.ExecutionPrice
	}

	return p.MainOrder.Price // if not executed return main order price
}

func FormatPrice(f float64, precision int) string {
	return strconv.FormatFloat(f, 'f', precision, 64)
}
