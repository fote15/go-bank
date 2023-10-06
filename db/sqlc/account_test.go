package db

import (
	"context"
	"database/sql"
	"github.com/stretchr/testify/require"
	"gobank/util"
	"testing"
	"time"
)

func CreateRandomAccount(t *testing.T) Account {
	arg := CreateAccountParams{
		Owner:    util.RandomOwner(),
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}
	account, err := testQuaries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, account.Owner, arg.Owner)
	require.Equal(t, account.Balance, arg.Balance)
	require.Equal(t, account.Currency, arg.Currency)

	require.NotZero(t, account.CreatedAt)
	require.NotZero(t, account.ID)
	return account
}
func TestCreateAccount(t *testing.T) {
	CreateRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	account1 := CreateRandomAccount(t)
	account2, err := testQuaries.GetAccount(context.Background(), account1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, account2)
	require.Equal(t, account2.ID, account1.ID)
	require.Equal(t, account2.Owner, account1.Owner)
	require.Equal(t, account2.Balance, account1.Balance)
	require.Equal(t, account2.CreatedAt, account1.CreatedAt)
	require.Equal(t, account2.Currency, account1.Currency)
	require.WithinDuration(t, account1.CreatedAt, account2.CreatedAt, time.Second)
}

func TestUpdateAccount(t *testing.T) {
	account1 := CreateRandomAccount(t)
	arg := UpdateAccountParams{
		ID:      account1.ID,
		Balance: util.RandomMoney(),
	}

	account2, err := testQuaries.UpdateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account2)

	require.Equal(t, account2.ID, account1.ID)
	require.Equal(t, account2.Owner, account1.Owner)
	require.Equal(t, account2.CreatedAt, account1.CreatedAt)
	require.Equal(t, account2.Currency, account1.Currency)
	require.WithinDuration(t, account1.CreatedAt, account2.CreatedAt, time.Second)
}

func TestDeleteAccount(t *testing.T) {
	account1 := CreateRandomAccount(t)

	err := testQuaries.DeleteAccount(context.Background(), account1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, account1)

	account, err := testQuaries.GetAccount(context.Background(), account1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, account)
}
