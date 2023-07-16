package wood

func ComponentLevel(label string, level Level) {
	componentLevel[label] = level
}

func PrefixLevel(label string, level Level) {
	prefixes.PutString(label, level)
}

//func splitName(name string) *logStack {
//	dot := strings.LastIndex(name, ".")
//	if dot == -1 {
//		return &logStack{
//			CanonicalID: name,
//			LastID:      name,
//			Display:     name,
//		}
//	}
//	return &logStack{
//		CanonicalID: name,
//		LastID:      name[dot+1:],
//		Display:     name[dot+1:],
//	}
//}

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
