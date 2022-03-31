package backtester

import (
	"context"
	"stregy/internal/domain/backtester"
	"stregy/test/quote"
	"testing"
	"time"
)

func NewMockedService() backtester.Service {
	repository := NewMockedRepository()
	return backtester.NewService(repository, quote.NewMockedService())
}

func TestRunBacktester(t *testing.T) {
	startDate, _ := time.Parse("2006-01-02", "2020-01-29")
	endDate := time.Now()
	bt := backtester.Backtester{
		StartDate: startDate,
		EndDate:   endDate,
	}

	service := NewMockedService()

	service.Run(context.TODO(), &bt)
}
