package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "todo-exercise/testutil/keeper"
	"todo-exercise/x/todo/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.TodoKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
