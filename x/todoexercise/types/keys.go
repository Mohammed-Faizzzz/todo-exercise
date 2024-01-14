package types

const (
	// ModuleName defines the module name
	ModuleName = "todoexercise"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_todoexercise"
)

var (
	ParamsKey = []byte("p_todoexercise")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
