package goconsole

import "fmt"

var colours = map[string]string{
	"reset":      "\033[0m",
	"bold":       "\033[1m",
	"dim":        "\033[2m",
	"italic":     "\033[3m",
	"underline":  "\033[4m",
	"blink":      "\033[5m",
	"invert":     "\033[7m",
	"hidden":     "\033[8m",

	"black":      "\033[30m",
	"red":        "\033[31m",
	"green":      "\033[32m",
	"yellow":     "\033[33m",
	"blue":       "\033[34m",
	"magenta":    "\033[35m",
	"cyan":       "\033[36m",
	"white":      "\033[37m",

	"brightBlack":   "\033[90m",
	"brightRed":     "\033[91m",
	"brightGreen":   "\033[92m",
	"brightYellow":  "\033[93m",
	"brightBlue":    "\033[94m",
	"brightMagenta": "\033[95m",
	"brightCyan":    "\033[96m",
	"brightWhite":   "\033[97m",
}

/*
PrintColour outputs coloured text to the terminal

This function formats the message using fmt.Sprintf and applies an ANSI colour code.
Use it the same way as fmt.Printf with colour names listed below

Examples:
func init() {
	PrintColour("blue", "Obi-Wan: Hello there\n")
	PrintColour("red", "Grievous: %v %v", "General", "Kenobi")
}

# Accepted Ansi Values:

# Base colours:
 | Name            | Value           |
 |-----------------|-----------------|
 | Black           | black           |
 | Red             | red             |
 | Green           | green           |
 | Yellow          | yellow          |
 | Blue            | blue            |
 | Magenta         | magenta         |
 | Cyan            | cyan            |
 | White           | white           |

# Bright colours:
 | Name            | Value           |
 |-----------------|-----------------|
 | Bright Black    | brightBlack     |
 | Bright Red      | brightRed       |
 | Bright Green    | brightGreen     |
 | Bright Yellow   | brightYellow    |
 | Bright Blue     | brightBlue      |
 | Bright Magenta  | brightMagenta   |
 | Bright Cyan     | brightCyan      |
 | Bright White    | brightWhite     |

# Text styles:
 | Name            | Value           |
 |-----------------|-----------------|
 | Reset           | reset           |
 | Bold            | bold            |
 | Dim             | dim             |
 | Italic          | italic          |
 | Underline       | underline       |
 | Blink           | blink           |
 | Invert          | invert          |
 | Hidden          | hidden          |
*/
func PrintColour(colour string, format string, args ...interface{}) {
	code, ok := colours[colour]
	if !ok {
		code = colours["reset"]
	}
	fmt.Printf("%s%s%s", code, fmt.Sprintf(format, args...), colours["reset"])
}