package kettle

import (
	"os"
	"strings"
)

func GetFileLines(path string) ([]string, error) {
	bs, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return strings.Split(string(bs), "\n"), nil
}
