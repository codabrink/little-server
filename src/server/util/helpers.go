package util
import (
	"fmt"
	"strings"
)

func P(a ...interface{}) {fmt.Println(a)}
func Concat(a ...string) string {
	var builder strings.Builder
	for _, b := range a {builder.WriteString(b)}
	return builder.String()
}
