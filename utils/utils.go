package utils

import (
	"fmt"
	"text/tabwriter"

	"github.com/jroimartin/gocui"
)

// Longest returns the length of the longest string in the given array
func Longest(strs []string) int {
	longest := 0
	for _, str := range strs {
		if len(str) > longest {
			longest = len(str)
		}
	}
	return longest
}

// Max returns the maximum of the two ints
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Min returns the minimum of the two ints
func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func PrintFigure(w *tabwriter.Writer, content string, v *gocui.View) {
	width, _ := v.Size()
	fmt.Fprintf(w, fmt.Sprintf("\u001b[33;1m%*s\u001b[0m\n", -width, fmt.Sprintf("%*s", (width+len(content))/2, content)))

	// uncentered
	// fmt.Fprintf(w, "\u001b[33m%s\u001b[0m\n", content)
}

func PrintName(w *tabwriter.Writer, content string, v *gocui.View) {
	width, _ := v.Size()
	fmt.Fprintf(w, fmt.Sprintf("\u001b[32;1m%*s\u001b[0m\n", -width, fmt.Sprintf("%*s", (width+len(content))/2, content)))

	// uncentered
	// fmt.Fprintf(w, "\u001b[32;1m%s\u001b[0m\n", content)
}

func BlackPrint(w *tabwriter.Writer, content interface{}, tabbed bool) {
	if tabbed {
		fmt.Fprintf(w, "\u001b[30m%s\u001b[0m\t", content)
	} else {
		fmt.Fprintf(w, "\u001b[30m%s\u001b[0m", content)
	}
}
