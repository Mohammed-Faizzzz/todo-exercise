package todo_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "todo-exercise/testutil/keeper"
	"todo-exercise/testutil/nullify"
	"todo-exercise/x/todo/module"
	"todo-exercise/x/todo/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		TaskList: []types.Task{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		TaskCount: 2,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.TodoKeeper(t)
	todo.InitGenesis(ctx, k, genesisState)
	got := todo.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.TaskList, got.TaskList)
	require.Equal(t, genesisState.TaskCount, got.TaskCount)
	// this line is used by starport scaffolding # genesis/test/assert
}
