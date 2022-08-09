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

func (r repository) GetByInterval(ctx context.Context, symbol string, startTime, endTime time.Time) ([]quote.Quote, error) {
	tableName := strings.ToLower(symbol) + "_m1_quotes"
	startTimeStr := utils.FormatTime(startTime)
	endTimeStr := utils.FormatTime(endTime)

	quotes := make([]Quote, 0)
	err := r.db.Table(tableName).Where("time >= ? AND time < ?", startTimeStr, endTimeStr).Find(&quotes).Error
	if err != nil {
		return nil, err
	}

	quotesDomain := make([]quote.Quote, 0, len(quotes))
	for _, q := range quotes {
		quotesDomain = append(quotesDomain, q.ToDomain())
	}
	return quotesDomain, err
}

func (r repository) Load(symbol, filePath, delimiter string) error {
	tableName := strings.ToLower(symbol) + "_M1_quotes"
	return r.db.Exec(fmt.Sprintf(`
	CREATE UNLOGGED TABLE IF NOT EXISTS temp_quotes (
		time double precision,
		open decimal(20, 8),
		high decimal(20, 8),
		low decimal(20, 8),
		close decimal(20, 8),
		volume real
	 );
	 
	COPY temp_quotes FROM '%v' DELIMITERS '%v' CSV;
	
	ALTER TABLE temp_quotes
	ALTER time TYPE timestamp without time zone
		USING (to_timestamp(time) AT TIME ZONE 'UTC');
	
	CREATE TABLE %v (LIKE temp_quotes INCLUDING ALL);

	INSERT INTO %v SELECT * FROM temp_quotes ON CONFLICT DO NOTHING;
	 
	DROP TABLE temp_quotes;`,
		filePath, delimiter, tableName, tableName)).Error
}
