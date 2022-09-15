package db

import (
	"context"
	"testing"

	"github.com/BaharaJr/go/util"
	"github.com/stretchr/testify/require"
)

func CreateRandomTransfer(t *testing.T) Transfer {
	sender := CreateRandomAccount(t)
	receiver := CreateRandomAccount(t)

	arg := CreateTransferParams{
		FromAccountID: sender.ID,
		ToAccountID:   receiver.ID,
		Amount:        util.RandomMoney(),
	}

	transfer, err := testQueries.CreateTransfer(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, transfer)
	require.Equal(t, sender.ID, transfer.FromAccountID)
	require.Equal(t, receiver.ID, transfer.ToAccountID)

	return transfer
}

func TestCreateTransfer(t *testing.T) {
	CreateRandomTransfer(t)
}

func TestGetTransfer(t *testing.T) {
	transfer := CreateRandomTransfer(t)

	getTransfer, err := testQueries.GetTransfer(context.Background(), transfer.ID)

	require.NoError(t, err)
	require.NotEmpty(t, getTransfer)
	require.Equal(t, getTransfer.FromAccountID, transfer.FromAccountID)
	require.Equal(t, getTransfer.ToAccountID, transfer.ToAccountID)
	require.Equal(t, getTransfer.ID, transfer.ID)
}
func TestListTransfers(t *testing.T) {
	arg := ListTransfersParams{
		Limit:  4,
		Offset: 0,
	}
	transfers, err := testQueries.ListTransfers(context.Background(), arg)

	require.NoError(t, err)
	require.Empty(t, transfers)
	require.Equal(t, len(transfers), 0)

	for _, transfer := range transfers {
		require.NotEmpty(t, transfer)
	}
}
