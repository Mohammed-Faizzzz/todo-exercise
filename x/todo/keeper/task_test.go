package keeper_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "todo-exercise/testutil/keeper"
	"todo-exercise/testutil/nullify"
	"todo-exercise/x/todo/keeper"
	"todo-exercise/x/todo/types"
)

func createNTask(keeper keeper.Keeper, ctx context.Context, n int) []types.Task {
	items := make([]types.Task, n)
	for i := range items {
		items[i].Id = keeper.AppendTask(ctx, items[i])
	}
	return items
}

func TestTaskGet(t *testing.T) {
	keeper, ctx := keepertest.TodoKeeper(t)
	items := createNTask(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetTask(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestTaskRemove(t *testing.T) {
	keeper, ctx := keepertest.TodoKeeper(t)
	items := createNTask(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveTask(ctx, item.Id)
		_, found := keeper.GetTask(ctx, item.Id)
		require.False(t, found)
	}
}

func TestTaskGetAll(t *testing.T) {
	keeper, ctx := keepertest.TodoKeeper(t)
	items := createNTask(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllTask(ctx)),
	)
}

func TestTaskCount(t *testing.T) {
	keeper, ctx := keepertest.TodoKeeper(t)
	items := createNTask(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetTaskCount(ctx))
}
