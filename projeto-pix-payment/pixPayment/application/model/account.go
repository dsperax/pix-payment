package model

import (
	"time"
	uuid "github.com/satori/go.uuid"
	"github.com/asaskevich/govalidator"
)

type Account struct {
	Base 		`valid:"required"`
	OwnerName   string `json:"owner_name" valid:"notnull"`
	Bank		*Bank `valid:"-"`
	Number 		string `json:"number" valid:"notnull"`
	PixKeys		[]*PixKey `valid:"-"`
}

func (account *Account) isValid() erro {
	_, err := govalidator.ValidateStruct(account)
	if err != nil {
		return err
	}
	return nil
}

func NewAccount(bank *Bank, number string, ownerName string) (*Account, error){
	account := Account{
		OwnerName: ownerName,
		Bank: bank,
		Number: number,
	}

	account.ID = uuid.NewV4().String()
	account.CreatedAt = time.Now()

	err := account.isValid()
	if err != nil {
		return nil, err
	}

	return &account, nil
}