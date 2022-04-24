package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/ramses2099/backendgo/util"
	"github.com/stretchr/testify/require"
)

func createRandomTransfers(t *testing.T) Transfer {

	limit := ListAccountsParams{
		Limit:  1,
		Offset: 0,
	}
	FromAccount, _ := testQueries.GetRandomAccountID(context.Background(), limit)
	ToAccountID, _ := testQueries.GetRandomAccountID(context.Background(), limit)

	arg := CreateTransferParams{
		FromAccount: FromAccount,
		ToAccountID: ToAccountID,
		Amount:      util.RandomMoney(),
	}

	transfers, err := testQueries.CreateTransfer(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, transfers)

	require.Equal(t, arg.FromAccount, transfers.FromAccount)
	require.Equal(t, arg.ToAccountID, transfers.ToAccountID)
	require.Equal(t, arg.Amount, transfers.Amount)

	require.NotZero(t, transfers.ID)
	require.NotZero(t, transfers.CreatedAt)

	return transfers
}

func TestCreateTransfer(t *testing.T) {
	createRandomTransfers(t)
}

func TestGetTransfer(t *testing.T) {
	transfer1 := createRandomTransfers(t)
	transfer2, err := testQueries.GetTransfer(context.Background(), transfer1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, transfer2)

	require.Equal(t, transfer1.ID, transfer2.ID)
	require.Equal(t, transfer1.FromAccount, transfer2.FromAccount)
	require.Equal(t, transfer1.ToAccountID, transfer2.ToAccountID)
	require.Equal(t, transfer1.Amount, transfer2.Amount)
	require.WithinDuration(t, transfer1.CreatedAt, transfer2.CreatedAt, time.Second)

}

func TestUpdateTranfer(t *testing.T) {
	transfer1 := createRandomTransfers(t)

	arg := UpdateTranferParams{
		ID:     transfer1.ID,
		Amount: util.RandomMoney(),
	}

	transfer2, err := testQueries.UpdateTranfer(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, transfer2)
	require.Equal(t, transfer1.ID, transfer2.ID)
	require.Equal(t, transfer1.FromAccount, transfer2.FromAccount)
	require.Equal(t, transfer1.ToAccountID, transfer2.ToAccountID)
	require.Equal(t, arg.Amount, arg.Amount)
	require.WithinDuration(t, transfer1.CreatedAt, transfer2.CreatedAt, time.Second)

}

func TestDeleteTranfer(t *testing.T) {
	transfer1 := createRandomTransfers(t)
	err := testQueries.DeleteTranfer(context.Background(), transfer1.ID)
	require.NoError(t, err)

	transfer2, err := testQueries.GetTransfer(context.Background(), transfer1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, transfer2)
}

func TestListTransfers(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomTransfers(t)
	}

	arg := ListTransfersParams{
		Limit:  5,
		Offset: 5,
	}

	entries, err := testQueries.ListTransfers(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, entries, 5)

	for _, entry := range entries {
		require.NotEmpty(t, entry)
	}
}
