package migration

import (
	"stregy/internal/adapters/pgorm/exchange"
	"stregy/internal/adapters/pgorm/exgaccount"
	"stregy/internal/adapters/pgorm/order"
	"stregy/internal/adapters/pgorm/position"
	"stregy/internal/adapters/pgorm/strategy"
	"stregy/internal/adapters/pgorm/stratexec"
	"stregy/internal/adapters/pgorm/symbol"
	"stregy/internal/adapters/pgorm/user"

	"gorm.io/gorm"
)

func createDatatypes(db *gorm.DB) error {
	return db.Exec(`DO $$ BEGIN CREATE TYPE order_type AS ENUM('LimitOrder', 'MarketOrder', 'StopLimitOrder', 'StopOrder', 'TrailingStopOrder', 'CloseByLimitOrder', 'CloseByStopOrder', 'CloseByMarketOrder'); EXCEPTION WHEN duplicate_object THEN null; END $$;
		DO $$ BEGIN CREATE TYPE order_status AS ENUM('SubmittedOrder', 'AcceptedOrder', 'RejectedOrder', 'PartialOrder', 'FilledOrder', 'CancelledOrder', 'ExpiredOrder', 'MarginOrder'); EXCEPTION WHEN duplicate_object THEN null; END $$;
		DO $$ BEGIN CREATE TYPE position_status AS ENUM('Draft', 'OpenPosition', 'TakeProfitPosition', 'StopLossPosition', 'MarketClosePosition'); EXCEPTION WHEN duplicate_object THEN null; END $$;
		DO $$ BEGIN CREATE TYPE order_direction AS ENUM('Long', 'Short'); EXCEPTION WHEN duplicate_object THEN null; END $$;
		DO $$ BEGIN CREATE TYPE strategy_execution_status AS ENUM('Created', 'Running', 'Finished', 'Crashed'); EXCEPTION WHEN duplicate_object THEN null; END $$;`,
	).Error
}

func createExtensions(db *gorm.DB) error {
	return db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`).Error
}

func Migrate(db *gorm.DB) error {
	err := createExtensions(db)
	if err != nil {
		return err
	}
	err = createDatatypes(db)
	if err != nil {
		return err
	}
	return db.AutoMigrate(&user.User{}, &symbol.Symbol{},
		&exchange.Exchange{}, &exgaccount.ExchangeAccount{},
		&strategy.Strategy{}, &stratexec.StrategyExecution{}, &order.Order{},
		&position.Position{})
}
