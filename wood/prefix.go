package wood

func ComponentLevel(level Level, label string) {
	componentLevel[label] = level
}

func PrefixLevel(level Level, label string) {
	prefixes.PutString(label, level)
}

func ignored(action Level) bool {
	if currentCanonical == "" {
		return false
	}

	// d is DEBUG
	// current is INFO (logger)
	// true if DEBUG(5) > INFO (4)
	if v, ok := componentLevel[currentCanonical]; ok {
		return action < v
	}

	// a.b.c.d is DEBUG
	// current is INFO (logger)
	// true if DEBUG(5) > INFO (4)
	if v, ok := prefixes.GetByString(currentCanonical); ok {
		return action < v
	}
	return false
}
