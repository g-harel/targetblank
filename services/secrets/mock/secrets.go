package mock

// RawKey exposes the mock key directly.
const RawKey = ""

// Key returns a hardcoded secret.
func Key() (string, error) {
	return RawKey, nil
}
