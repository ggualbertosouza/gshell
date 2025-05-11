package prompt

type Color string

const (
	Red     Color = "red"
	Green   Color = "green"
	Yellow  Color = "yellow"
	Blue    Color = "blue"
	Magenta Color = "magenta"
	Cyan    Color = "cyan"
	Reset   Color = "reset"
)

var colorCodes = map[Color]string{
	Red:     "\033[31m",
	Green:   "\033[32m",
	Yellow:  "\033[33m",
	Blue:    "\033[34m",
	Magenta: "\033[35m",
	Cyan:    "\033[36m",
	Reset:   "\033[0m",
}

func Print(text string, color Color) string {
	code, ok := colorCodes[color]
	if !ok {
		code = colorCodes[Reset]
	}

	return code + text + colorCodes[Reset]
}
