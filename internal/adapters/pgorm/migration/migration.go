package migration

import (
	"stregy/internal/adapters/pgorm/exchange"
	"stregy/internal/adapters/pgorm/exgaccount"
	"stregy/internal/adapters/pgorm/order"
	"stregy/internal/adapters/pgorm/position"
	"stregy/internal/adapters/pgorm/quote"
	"stregy/internal/adapters/pgorm/stratexec"
	"stregy/internal/adapters/pgorm/symbol"
	"stregy/internal/adapters/pgorm/tick"

	"gorm.io/gorm"
)

func createDatatypes(db *gorm.DB) error {
	return db.Exec(`DO $$ BEGIN CREATE TYPE order_type AS ENUM('LimitOrder', 'MarketOrder', 'StopLimitOrder', 'StopMarketOrder', 'TrailingStopOrder'); EXCEPTION WHEN duplicate_object THEN null; END $$;
		DO $$ BEGIN CREATE TYPE order_status AS ENUM('SubmittedOrder', 'AcceptedOrder', 'PartialOrder', 'CompletedOrder', 'CancelledOrder', 'ExpiredOrder', 'MarginOrder'); EXCEPTION WHEN duplicate_object THEN null; END $$;
		DO $$ BEGIN CREATE TYPE position_status AS ENUM('CreatedPosition', 'PartialPosition', 'OpenPosition', 'TakeProfitPosition', 'StopLossPosition', 'CancelledPosition'); EXCEPTION WHEN duplicate_object THEN null; END $$;
		DO $$ BEGIN CREATE TYPE order_direction AS ENUM('Long', 'Short'); EXCEPTION WHEN duplicate_object THEN null; END $$;
		DO $$ BEGIN CREATE TYPE strategy_execution_status AS ENUM('Created', 'Running', 'Finished', 'Crashed'); EXCEPTION WHEN duplicate_object THEN null; END $$;`,
	).Error
}

func createExtensions(db *gorm.DB) error {
	return db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`).Error
}

func Migrate(db *gorm.DB) error {
	if err := createExtensions(db); err != nil {
		return err
	}
	if err := createDatatypes(db); err != nil {
		return err
	}

	return db.AutoMigrate(&symbol.Symbol{}, &tick.Tick{},
		&quote.Quote{}, &exchange.Exchange{}, &exgaccount.ExchangeAccount{},
		&stratexec.StrategyExecution{}, &position.Position{}, &order.Order{})
}
