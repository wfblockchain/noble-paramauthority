package proposal

const (
	AuthorityKey = "authority"

	ModuleName = "paramauthority"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
