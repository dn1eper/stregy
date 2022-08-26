package client

// import (
// 	"context"
// 	"stregy/internal/composites"
// 	"stregy/internal/config"

// 	log "github.com/sirupsen/logrus"
// )

// func CreateClientTestObjects() {
// 	cfg := config.GetConfig()

// 	pgormComposite, err := composites.NewPGormComposite(context.Background(), cfg.PosgreSQL.Host, cfg.PosgreSQL.Port, cfg.PosgreSQL.Username, cfg.PosgreSQL.Password, cfg.PosgreSQL.Database)
// 	if err != nil {
// 		log.Fatal("pgorm composite failed")
// 	}

// 	userComposite, err := composites.NewUserComposite(pgormComposite)
// 	if err != nil {
// 		log.Fatal("user composite failed")
// 	}

// 	exgAccountComposite, err := composites.NewExchangeAccountComposite(pgormComposite, userComposite.Service)
// 	if err != nil {
// 		log.Fatal("exchange account composite failed")
// 	}

// }
