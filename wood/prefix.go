package wood

import "strings"

func PrefixLevel(label string, level Level) {
	prefixes.PutString(label, level)
}

type logStack struct {
	Id      string
	Display string
}

func splitName(name string) *logStack {
	dot := strings.LastIndex(name, ".")
	if dot == -1 {
		return &logStack{
			Id:      "",
			Display: name,
		}
	}
	return &logStack{
		Id:      name,
		Display: name[dot+1:],
	}
}

func ignored(action Level) bool {
	if currentId == "" {
		return false
	}

	v, ok := prefixes.GetByString(currentId)
	if !ok {
		return false
	}
	return v < action
}
