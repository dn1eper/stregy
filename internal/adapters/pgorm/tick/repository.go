package tick

import (
	"fmt"
	"stregy/internal/domain/tick"
	"stregy/pkg/utils"
	"strings"
	"time"
	"unsafe"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(client *gorm.DB) tick.Repository {
	return &repository{db: client}
}

func (r repository) GetByInterval(symbol string, startTime, endTime time.Time) ([]tick.Tick, error) {
	tableName := strings.ToLower(symbol) + "_ticks"
	startTimeStr := utils.FormatTime(startTime)
	endTimeStr := utils.FormatTime(endTime)

	ticks := make([]Tick, 0)
	err := r.db.Table(tableName).Where("time >= ? AND time < ?", startTimeStr, endTimeStr).Find(&ticks).Error
	if err != nil {
		return nil, err
	}

	_ = tick.Tick(Tick{}) // check Tick and tick.Tick have same fields
	ticksDomain := (*[]tick.Tick)(unsafe.Pointer(&ticks))
	return *ticksDomain, nil
}

func (r repository) Load(symbol, filePath, delimiter string) error {
	tableName := strings.ToLower(symbol) + "_ticks"
	return r.db.Exec(fmt.Sprintf(`
	CREATE UNLOGGED TABLE IF NOT EXISTS temp_ticks (
		time double precision,
		price decimal(20, 8)
	 );
	 
	COPY temp_ticks FROM '%v' DELIMITERS '%v' CSV;
	
	ALTER TABLE temp_ticks
	ALTER time TYPE timestamp without time zone
		USING (to_timestamp(time) AT TIME ZONE 'UTC');
	 
	CREATE TABLE %v (LIKE temp_quotes INCLUDING ALL);

	INSERT INTO %v SELECT * FROM temp_ticks ON CONFLICT DO NOTHING;

	DROP TABLE temp_ticks;`,
		filePath, delimiter, tableName, tableName)).Error
}
