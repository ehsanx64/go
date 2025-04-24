package general

import (
	"fmt"
	"strings"
)

const defaultColumnWidth = 15

var columnWidth int = defaultColumnWidth

func SetColumnWidth(width int) {
	columnWidth = width
}

func DumpValue(title string, arg any) {
	var padding strings.Builder

	for i := 0; i < columnWidth-(len(title)+1); i++ {
		padding.WriteString(" ")
	}

	fmt.Printf("%s:%s%+v\n", title, padding.String(), arg)
}

func DumpType(title string, arg any) {
	var padding strings.Builder

	for i := 0; i < columnWidth-(len(title)+1); i++ {
		padding.WriteString(" ")
	}

	fmt.Printf("%s:%s%+T\n", title, padding.String(), arg)
}
