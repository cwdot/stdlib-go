package wood

func PrefixLevel(level Level, label string) {
	prefixes.PutString(label, level)
}

func ignored(action Level) bool {
	if currentCanonical == "" {
		return false
	}

	show := false
	for _, p := range stack {
		// a.b.c.d is DEBUG
		// current is INFO (logger)
		// true if DEBUG(5) > INFO (4)
		if v, ok := prefixes.GetByString(p.CanonicalID); ok {
			show = action < v
		}
	}
	return show
}
