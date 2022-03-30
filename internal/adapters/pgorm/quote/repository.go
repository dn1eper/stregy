package quote

import (
	"context"
	"fmt"
	"stregy/internal/domain/quote"
	"stregy/pkg/utils"
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

func (r *repository) GetByInterval(ctx context.Context, symbol string, startTime, endTime time.Time, offset, pageSize int) ([]quote.Quote, error) {
	tableName := strings.ToLower(symbol) + "s"
	startTimeStr := utils.FormatTime(startTime)
	endTimeStr := utils.FormatTime(endTime)

	quotes := make([]Quote, 0)
	err := r.db.Table(tableName).Offset(offset).Limit(pageSize).Where("time >= ? AND time <= ?", startTimeStr, endTimeStr).Find(&quotes).Error
	if err != nil {
		return nil, err
	}

	quotesDomain := make([]quote.Quote, 0, len(quotes))
	for _, q := range quotes {
		quotesDomain = append(quotesDomain, q.ToDomain())
	}
	return quotesDomain, err
}

func (r *repository) Load(ctx context.Context, symbol, filePath, delimiter string) error {
	tableName := strings.ToLower(symbol) + "s"
	return r.db.Exec(fmt.Sprintf("COPY %v FROM '%v' DELIMITERS '%v' CSV;", tableName, filePath, delimiter)).Error
}
