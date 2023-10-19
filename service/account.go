package service

import (
	"context"
	"fmt"
	"github.com/bnb-chain/mind-marketplace-backend/dao"
	"github.com/bnb-chain/mind-marketplace-backend/models"

	"gorm.io/gorm"
)

type Account interface {
	Get(context context.Context, accountName string) (*models.Account, error)
}

type AccountService struct {
	accountDao dao.AccountDao
}

func NewAccountService(accountDao dao.AccountDao) Account {
	return &AccountService{
		accountDao: accountDao,
	}
}

func (s *AccountService) Get(context context.Context, address string) (*models.Account, error) {
	if err := validateAddress(address); err != nil {
		return nil, err
	}

	account, err := s.accountDao.GetByAddress(context, address)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, NotFoundErr
		} else {
			return nil, fmt.Errorf("fail to read account from database")
		}
	}

	return convertAccount(account), nil
}
