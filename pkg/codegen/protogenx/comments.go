package protogenx

import (
	"strings"

	"github.com/RyoJerryYu/go-utilx/pkg/container/slicex"
)

type Comments string

func (c Comments) String() string {
	lines := strings.Split(string(c), "\n")
	lines = slicex.To(lines, func(line string) string {
		return "// " + line
	})
	return strings.Join(lines, "\n")
}
