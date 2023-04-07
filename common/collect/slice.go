package collect

func Contain[T comparable](sl []T, t T) bool {
	for _, v := range sl {
		if v == t {
			return true
		}
	}
	return false
}
