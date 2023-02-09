package quote

import (
	"fmt"
	"sort"
	"stregy/internal/domain/quote"
	"stregy/pkg/utils"
	"strings"
	"time"
	"unsafe"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(client *gorm.DB) quote.Repository {
	return &repository{db: client}
}

func (r repository) Get(symbol string, startTime, endTime time.Time, limit, timeframeSec int) ([]quote.Quote, error) {
	tableName := getTableName(symbol, timeframeSec)
	startTimeStr := utils.FormatTime(startTime)
	endTimeStr := utils.FormatTime(endTime)

	quotes := make([]Quote, 0)
	err := r.db.Table(tableName).Where("time >= ? AND time <= ? ORDER BY time LIMIT ?", startTimeStr, endTimeStr, limit).Find(&quotes).Error
	if err != nil {
		return nil, err
	}

	res := *(*[]quote.Quote)(unsafe.Pointer(&quotes))
	if res[len(res)-1].Time.After(endTime) {
		i := sort.Search(len(res), func(i int) bool {
			return res[i].Time.After(endTime)
		})
		res = res[:i]
	}

	return res, nil
}

func (r repository) Load(symbol, filePath, delimiter string, timeframeSec int) error {
	tableName := getTableName(symbol, timeframeSec)
	return r.db.Exec(fmt.Sprintf(`
	CREATE UNLOGGED TABLE IF NOT EXISTS temp_quotes (
		time double precision,
		open double precision,
		high double precision,
		low double precision,
		close double precision,
		volume int
	 );
	 
	COPY temp_quotes FROM '%v' DELIMITERS '%v' CSV;
	
	ALTER TABLE temp_quotes
	ALTER time TYPE timestamp without time zone
		USING (to_timestamp(time) AT TIME ZONE 'UTC');
	
	CREATE TABLE IF NOT EXISTS %v (LIKE quotes INCLUDING ALL);

	INSERT INTO %v SELECT * FROM temp_quotes ON CONFLICT DO NOTHING;
	 
	DROP TABLE temp_quotes;`,
		filePath, delimiter, tableName, tableName)).Error
}

func getTableName(symbol string, timeframeSec int) string {
	tableName := strings.ToLower(symbol)
	if timeframeSec < 60 {
		tableName += "_s1_quotes"
	} else {
		tableName += "_m1_quotes"
	}

	return tableName
}
