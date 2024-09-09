package wood

func PrefixLevel(level Level, label string) {
	prefixes.PutString(label, level)
}

func ComponentLevel(level Level, label string) {
	components[label] = level
}

func ignored(action Level) bool {
	if currentCanonical == "" {
		return false
	}

	ignore := false // we need to get the most specific one (must complete the loop)
	for _, p := range stack {
		// a.b.c.d is DEBUG
		// current is INFO (logger)
		// true if DEBUG(5) > INFO (4)
		if v, ok := components[p.LastID]; ok {
			ignore = action > v
		} else if v, ok := prefixes.GetByString(p.CanonicalID); ok {
			ignore = action > v
		}
		if ignore {
			break
		}
	}
	return ignore
}
