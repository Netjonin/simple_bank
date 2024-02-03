package db

import (
	"context"
	"testing"
	"time"

	"github.com/netjonin/simplebank/util"
	"github.com/stretchr/testify/require"
)

func createRandomEntry(t *testing.T) Entry {
	Ids, err := testQueries.ListAccountIDs(context.Background())
	require.NoError(t, err)

	Id := util.RandomInt(Ids[0], Ids[len(Ids) - 1])
	arg := CreateEntryParams {
		AccountID: Id,
		Amount: util.RandomMoney(),
	}

	entry, err := testQueries.CreateEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entry)
	require.Equal(t, arg.AccountID, entry.AccountID)
	require.Equal(t, arg.Amount, entry.Amount)

	require.NotZero(t, entry.AccountID)
	require.NotZero(t, entry.CreatedAt)

	return entry
}

func TestCreateEntry(t *testing.T) {
	createRandomEntry(t)
}

