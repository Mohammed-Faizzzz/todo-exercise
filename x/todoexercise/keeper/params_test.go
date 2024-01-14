package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "todo-exercise/testutil/keeper"
	"todo-exercise/x/todoexercise/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.TodoexerciseKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
