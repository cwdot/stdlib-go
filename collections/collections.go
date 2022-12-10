package collections

func Contains[T comparable](s []T, term T) bool {
	for _, v := range s {
		if v == term {
			return true
		}
	}
	return false
}
