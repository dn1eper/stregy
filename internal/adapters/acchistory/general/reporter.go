package acchistory

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

/*
resulting csv format:

	order id, position id, contingent type, diraction, type, size, submision time, submision price, execution time, execution price
*/
func (accountHistoryReporter) CreateReport(orders []*order.Order, s symbol.Symbol, filePath string) error {
	f, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer f.Close()

	for _, o := range orders {
		f.WriteString(fmt.Sprintf(
			"%d,%d,%s,%s,%s,%f,%s,%s,%s,%s\n",
			o.ID,
			o.Position.ID,
			getOrderContingentType(o),
			o.Diraction.String(),
			o.Type.String(),
			o.Size,
			o.SubmissionTime.Format("2006-01-02 15:04:05"),
			FormatPrice(o.Price, s.Precision),
			o.FCTime.Format("2006-01-02 15:04:05"),
			FormatPrice(o.ExecutionPrice, s.Precision)))
	}

	return nil
}

func getOrderContingentType(o *order.Order) string {
	p := o.Position
	if p.MainOrder.ID == o.ID {
		return "Main"
	}
	if p.StopOrder.ID == o.ID {
		return "SL"
	}
	if p.TakeOrder.ID == o.ID {
		return "TP"
	}

	panic("order does not belong to its position")
}

func FormatPrice(f float64, precision int) string {
	return strconv.FormatFloat(f, 'f', precision, 64)
}
