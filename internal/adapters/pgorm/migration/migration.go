package migration

import (
	"stregy/internal/adapters/pgorm/exchange"
	"stregy/internal/adapters/pgorm/exgaccount"
	"stregy/internal/adapters/pgorm/order"
	"stregy/internal/adapters/pgorm/position"
	"stregy/internal/adapters/pgorm/stratexec"
	"stregy/internal/adapters/pgorm/symbol"
	"stregy/internal/adapters/pgorm/user"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(&user.User{}, &symbol.Symbol{}, &exchange.Exchange{},
		&exgaccount.ExchangeAccount{}, &stratexec.StrategyExecution{},
		&order.Order{}, &position.Position{})
}
