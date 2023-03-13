package db

import (
	"context"
	"github.com/ku-kim/learn-go-with-project-simplebank/util"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestCreateEntry(t *testing.T) {
	account := createRandomAccount(t)
	createRandomEntry(t, account)
}

func createRandomEntry(t *testing.T, account Account) Entry {
	arg := CreateEntryParams{
		AccountID: account.ID,
		Amount:    util.RandomMoney(),
	}

	entry, err := testQueries.CreateEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.Equal(t, arg.AccountID, entry.AccountID)
	require.Equal(t, arg.Amount, entry.Amount)

	require.NotZero(t, entry.ID)
	require.NotZero(t, entry.CreatedAt)

	return entry
}

func TestGetEntry(t *testing.T) {
	account := createRandomAccount(t)
	createdEntry := createRandomEntry(t, account)
	entry, err := testQueries.GetEntry(context.Background(), createdEntry.ID)
	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.Equal(t, entry.ID, createdEntry.ID)
	require.Equal(t, entry.AccountID, createdEntry.AccountID)
	require.Equal(t, entry.Amount, createdEntry.Amount)
	require.WithinDuration(t, entry.CreatedAt, createdEntry.CreatedAt, time.Second)
}

func TestListEntries(t *testing.T) {
	account := createRandomAccount(t)
	for i := 0; i < 10; i++ {
		createRandomEntry(t, account)
	}

	arg := ListEntriesParams{
		AccountID: account.ID,
		Limit:     5,
		Offset:    5,
	}

	entries, err := testQueries.ListEntries(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, entries, 5)

	for _, entry := range entries {
		require.NotEmpty(t, entry)
		require.Equal(t, arg.AccountID, entry.AccountID)
	}
}
