package model

import (
	"time"
	uuid "github.com/satori/go.uuid"
	"github.com/asaskevich/govalidator"
)

type PixKeyRepositoryInterface interface {
	RegisterKey(pixKey *PixKey) (*PixKey, error)
	FindKeyByFind(key string, kind string) (*PixKey, error)
	AddBank(bank *Bank) error
	AddAccount(account *Account) error
	FindAccount(id string) (*Account, error)
}

type PixKey struct {
	Base 		`valid:"required"`
	Kind 		string `json:"kind" valid:"notnull"`
	Key 		string `json:"key" valid:"notnull"`
	AccountID 	string `json:"account_id" valid:"notnull"`
	Account 	*Account `valid:"-"`
	Status 		string `json:"status" valid:"notnull"`
}

func (pixKey *PixKey) isValid() error {
	-, err := govalidator.ValidateStruct(pixKey)

	if pixKey.Kind != "email" && pixKey.Kind != "cpf" {
		return errors.New(text: "Invalid type of key")
	}

	if pixKey.Status != "active" && pixKey.Status != "inactive" {
		return errors.New(text: "Invalid status")
	}

	if err != nil {
		return err
	}
	return nil
}

func NewPixKey(kind string, account *Account, key string) (*PixKey, error){
	pixKey := PixKey{
		Kind: kind,
		Key: key,
		Account: account,
		Status: "active",
	}

	pixKey.ID = uuid.NewV4().String()
	pixKey.CreatedAt = time.Now()

	err := pixKey.isValid()
	if err != nil {
		return nil, err
	}

	return &pixKey, nil
}