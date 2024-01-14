package todoexercise_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "todo-exercise/testutil/keeper"
	"todo-exercise/testutil/nullify"
	"todo-exercise/x/todoexercise/module"
	"todo-exercise/x/todoexercise/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.TodoexerciseKeeper(t)
	todoexercise.InitGenesis(ctx, k, genesisState)
	got := todoexercise.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
