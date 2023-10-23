package service

import (
	"context"
	"fmt"
	"github.com/bnb-chain/mind-marketplace-backend/dao"
	"github.com/bnb-chain/mind-marketplace-backend/database"
	"github.com/bnb-chain/mind-marketplace-backend/models"
	"time"

	"gorm.io/gorm"
)

type Account interface {
	Get(context context.Context, address string) (*models.Account, error)
	Update(context context.Context, request *models.UpdateAccountRequest) (*models.Account, error)
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

func (s *AccountService) Update(context context.Context, request *models.UpdateAccountRequest) (*models.Account, error) {
	//check timestamp
	current := time.Now().Unix()
	if current < *request.Timestamp || (current-*request.Timestamp) > 600 {
		return nil, InvalidTimestampErr
	}

	//check existence
	needCreate := false
	existed, err := s.accountDao.GetByAddress(context, *request.Address)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			needCreate = true
		} else {
			return nil, fmt.Errorf("fail to read account from database")
		}
	}

	//todo: verify signature

	if needCreate {
		bio := ""
		if request.Bio != nil {
			bio = *request.Bio
		}
		//save to database
		record := database.Account{
			Address:           *request.Address,
			Bio:               bio,
			TwitterUserName:   request.TwitterUserName,
			InstagramUserName: request.InstagramUserName,
		}

		if err := s.accountDao.Create(context, &record); err != nil {
			return nil, fmt.Errorf("fail to create account")
		}
		existed = record
	} else {
		//only update allowing fields
		if request.Bio != nil && *request.Bio != "" {
			existed.Bio = *request.Bio
		}
		if request.TwitterUserName != "" {
			existed.TwitterUserName = request.TwitterUserName
		}
		if request.InstagramUserName != "" {
			existed.InstagramUserName = request.InstagramUserName
		}

		if err := s.accountDao.Update(context, &existed); err != nil {
			return nil, fmt.Errorf("fail to update account")
		}
	}
	return convertAccount(existed), nil
}
