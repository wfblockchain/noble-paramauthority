package proposal

const (
	AuthorityKey = "authority"

	ModuleName = "paramauthority"

	StoreKey = "paramauthority"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
