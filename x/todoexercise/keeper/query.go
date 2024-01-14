package keeper

import (
	"todo-exercise/x/todoexercise/types"
)

var _ types.QueryServer = Keeper{}
