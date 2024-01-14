package keeper

import (
	"todo-exercise/x/todo/types"
)

var _ types.QueryServer = Keeper{}
