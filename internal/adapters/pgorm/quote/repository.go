package quote

import (
	"context"
	"fmt"
	"stregy/internal/domain/quote"
	"strings"
	"time"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(client *gorm.DB) quote.Repository {
	return &repository{db: client}
}

func (r repository) GetByIntervalPaginate(ctx context.Context, symbol string, startTime, endTime time.Time, offset, pageSize int) ([]quote.Quote, error) {
	tableName := strings.ToLower(symbol) + "s"
	startTimeStr := fmt.Sprintf("%v", startTime)
	startTimeStr = startTimeStr[:len(startTimeStr)-4]
	endTimeStr := fmt.Sprintf("%v", endTime)
	endTimeStr = endTimeStr[:len(endTimeStr)-4]

	quotes := make([]Quote, 0)
	err := r.db.Table(tableName).Offset(offset).Limit(pageSize).Where("time >= ? AND time <= ?", startTimeStr, endTimeStr).Find(&quotes).Error
	if err != nil {
		return nil, err
	}

	quotesDomain := make([]quote.Quote, 0, len(quotes))
	for _, q := range quotes {
		quotesDomain = append(quotesDomain, quote.Quote{Time: q.Time, Open: q.Open, High: q.High, Low: q.Low, Close: q.Close})
	}
	return quotesDomain, err
}

func (r repository) Load(ctx context.Context, symbol, filePath, delimiter string) error {
	tableName := strings.ToLower(symbol) + "s"
	tableExistsQuery := fmt.Sprintf(`SELECT * FROM %v LIMIT 1;`, tableName)
	err := r.db.Exec(tableExistsQuery).Error
	if err != nil {
		// Create table.
		err = r.db.Exec(fmt.Sprintf("CREATE TABLE %v (LIKE quotes INCLUDING ALL);", tableName)).Error
		if err != nil {
			return err
		}
	}
	return r.db.Exec(fmt.Sprintf("COPY %v FROM '%v' DELIMITERS '%v' CSV;", tableName, filePath, delimiter)).Error
}
