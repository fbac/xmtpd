package indexer

import (
	"context"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/xmtp/xmtpd/pkg/db/queries"
	"github.com/xmtp/xmtpd/pkg/testutils"
)

const CONTRACT_ADDRESS = "0x0000000000000000000000000000000000000000"

func TestInitialize(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	db, _, cleanup := testutils.NewDB(t, ctx)
	defer cleanup()
	querier := queries.New(db)

	tracker, err := NewBlockTracker(ctx, CONTRACT_ADDRESS, querier)
	require.NoError(t, err)
	require.NotNil(t, tracker)
	require.Equal(t, uint64(0), tracker.GetLatestBlock())
}

func TestUpdateLatestBlock(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	db, _, cleanup := testutils.NewDB(t, ctx)
	defer cleanup()
	querier := queries.New(db)

	tracker, err := NewBlockTracker(ctx, CONTRACT_ADDRESS, querier)
	require.NoError(t, err)

	// Test updating to a higher block
	err = tracker.UpdateLatestBlock(ctx, 100)
	require.NoError(t, err)
	require.Equal(t, uint64(100), tracker.GetLatestBlock())

	// Test updating to a lower block (should not update)
	err = tracker.UpdateLatestBlock(ctx, 50)
	require.NoError(t, err)
	require.Equal(t, uint64(100), tracker.GetLatestBlock())

	// Test updating to the same block (should not update)
	err = tracker.UpdateLatestBlock(ctx, 100)
	require.NoError(t, err)
	require.Equal(t, uint64(100), tracker.GetLatestBlock())

	// Verify persistence
	newTracker, err := NewBlockTracker(ctx, CONTRACT_ADDRESS, querier)
	require.NoError(t, err)
	require.Equal(t, uint64(100), newTracker.GetLatestBlock())
}

func TestConcurrentUpdates(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	db, _, cleanup := testutils.NewDB(t, ctx)
	defer cleanup()
	querier := queries.New(db)

	tracker, err := NewBlockTracker(ctx, CONTRACT_ADDRESS, querier)
	require.NoError(t, err)

	var wg sync.WaitGroup
	numGoroutines := 10
	updatesPerGoroutine := 100

	// Launch multiple goroutines to update the block number
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(startBlock int) {
			defer wg.Done()
			for j := 0; j < updatesPerGoroutine; j++ {
				blockNum := uint64(startBlock + j)
				err := tracker.UpdateLatestBlock(ctx, blockNum)
				require.NoError(t, err)
			}
		}(i * updatesPerGoroutine)
	}

	wg.Wait()

	// The final block number should be the highest one attempted
	expectedFinalBlock := uint64((numGoroutines-1)*updatesPerGoroutine + (updatesPerGoroutine - 1))
	require.Equal(t, expectedFinalBlock, tracker.GetLatestBlock())

	// Verify persistence
	newTracker, err := NewBlockTracker(ctx, CONTRACT_ADDRESS, querier)
	require.NoError(t, err)
	require.Equal(t, expectedFinalBlock, newTracker.GetLatestBlock())
}

func TestMultipleContractAddresses(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	db, _, cleanup := testutils.NewDB(t, ctx)
	defer cleanup()
	querier := queries.New(db)

	address1 := "0x0000000000000000000000000000000000000001"
	address2 := "0x0000000000000000000000000000000000000002"

	tracker1, err := NewBlockTracker(ctx, address1, querier)
	require.NoError(t, err)
	tracker2, err := NewBlockTracker(ctx, address2, querier)
	require.NoError(t, err)

	// Update trackers independently
	err = tracker1.UpdateLatestBlock(ctx, 100)
	require.NoError(t, err)
	err = tracker2.UpdateLatestBlock(ctx, 200)
	require.NoError(t, err)

	// Verify different addresses maintain separate block numbers
	require.Equal(t, uint64(100), tracker1.GetLatestBlock())
	require.Equal(t, uint64(200), tracker2.GetLatestBlock())

	// Verify persistence for both addresses
	newTracker1, err := NewBlockTracker(ctx, address1, querier)
	require.NoError(t, err)
	newTracker2, err := NewBlockTracker(ctx, address2, querier)
	require.NoError(t, err)

	require.Equal(t, uint64(100), newTracker1.GetLatestBlock())
	require.Equal(t, uint64(200), newTracker2.GetLatestBlock())
}
