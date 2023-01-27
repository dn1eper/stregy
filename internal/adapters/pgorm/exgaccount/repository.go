package exgaccount

import (
	"fmt"
	"stregy/internal/domain/exgaccount"
	user1 "stregy/internal/domain/user"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(client *gorm.DB) exgaccount.Repository {
	return &repository{db: client}
}

func (r *repository) Create(exgAccount exgaccount.ExchangeAccount, user *user1.User) (*exgaccount.ExchangeAccount, error) {
	userID, err := uuid.Parse(user.ID)
	exgAccountDB := &ExchangeAccount{
		UserID:              userID,
		ConnectionString:    exgAccount.ConnectionString,
		ExchangeAccountName: exgAccount.ExchangeAccountName,
	}
	if exgAccount.ExchangeID != "" {
		exgUUID, _ := uuid.Parse(exgAccount.ExchangeID)
		fmt.Printf("parsed exgUUID = %v\n", exgUUID)
		exgAccountDB.ExchangeID = &exgUUID
	}

	err = r.db.Create(exgAccountDB).Error
	if err != nil {
		return nil, err
	}
	return exgAccountDB.ToDomain(), nil
}

func (r *repository) GetAll(userID string) ([]*exgaccount.ExchangeAccount, error) {
	userUUID, err := uuid.Parse(userID)
	if err != nil {
		return nil, err
	}
	exgAccount := &ExchangeAccount{
		UserID: userUUID,
	}
	var result []*ExchangeAccount
	r.db.Find(exgAccount, result)

	exgAccountsDomain := make([]*exgaccount.ExchangeAccount, len(result))
	for _, exgAccount := range result {
		exgAccountsDomain = append(exgAccountsDomain, exgAccount.ToDomain())
	}
	return exgAccountsDomain, nil
}

func (r *repository) GetOne(exgAccountID string) (*exgaccount.ExchangeAccount, error) {
	uuid, _ := uuid.Parse(exgAccountID)
	exgAccount := ExchangeAccount{ExchangeAccountID: uuid}
	err := r.db.First(&exgAccount).Error
	if err != nil {
		return nil, err
	}
	return exgAccount.ToDomain(), nil
}
