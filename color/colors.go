package color

import "fmt"

type Color string

func It(c Color, text string) string {
	return fmt.Sprintf("%s%s%s", c, text, Reset)
}

const (
	Normal       Color = ""
	Reset        Color = "\033[m"
	Bold         Color = "\033[1m"
	Red          Color = "\033[31m"
	Green        Color = "\033[32m"
	Yellow       Color = "\033[33m"
	Blue         Color = "\033[34m"
	Magenta      Color = "\033[35m"
	Cyan         Color = "\033[36m"
	BoldRed      Color = "\033[1;31m"
	BoldGreen    Color = "\033[1;32m"
	BoldYellow   Color = "\033[1;33m"
	BoldBlue     Color = "\033[1;34m"
	BoldMagenta  Color = "\033[1;35m"
	BoldCyan     Color = "\033[1;36m"
	FaintRed     Color = "\033[2;31m"
	FaintGreen   Color = "\033[2;32m"
	FaintYellow  Color = "\033[2;33m"
	FaintBlue    Color = "\033[2;34m"
	FaintMagenta       = "\033[2;35m"
	FaintCyan    Color = "\033[2;36m"
	BgRed        Color = "\033[41m"
	BgGreen      Color = "\033[42m"
	BgYellow     Color = "\033[43m"
	BgBlue       Color = "\033[44m"
	BgMagenta    Color = "\033[45m"
	BgCyan       Color = "\033[46m"
	Faint        Color = "\033[2m"
	FaintItalic  Color = "\033[2;3m"
	Reverse      Color = "\033[7m"
)
