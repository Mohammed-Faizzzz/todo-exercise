package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "todo-exercise/testutil/keeper"
	"todo-exercise/x/todo/types"
)

func TestParamsQuery(t *testing.T) {
	keeper, ctx := keepertest.TodoKeeper(t)
	params := types.DefaultParams()
	require.NoError(t, keeper.SetParams(ctx, params))

	response, err := keeper.Params(ctx, &types.QueryParamsRequest{})
	require.NoError(t, err)
	require.Equal(t, &types.QueryParamsResponse{Params: params}, response)
}
