package model_test

import (
	"testing"

	uuid "github.com/satori/go.uuid"

	"github.com/dsperax/pix-payment/projeto-pix-payment/pixPayment/domain/model"
	"github.com/stretchr/testify/require"
)

func TestModel_NewBank(t *testing.T) {

	code := "001"
	name := "Goliath N. Bank"

	bank, err := model.NewBank(code, name)

	require.Nil(t, err)
	require.NotEmpty(t, uuid.FromStringOrNil(bank.ID))
	require.Equal(t, bank.Code, code)
	require.Equal(t, bank.Name, name)

	_, err = model.NewBank("", "")
	require.NotNil(t, err)
}
