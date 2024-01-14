package types

const (
	// ModuleName defines the module name
	ModuleName = "todo"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_todo"
)

var (
	ParamsKey = []byte("p_todo")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	TaskKey      = "Task/value/"
	TaskCountKey = "Task/count/"
)
