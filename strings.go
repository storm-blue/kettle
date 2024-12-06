package kettle

import (
	"strings"
)

func ToLines(content string) []string {
	return strings.Split(content, "\n")
}
