package norwegish

import "strings"

func Translate(s string) string {
	r := strings.NewReplacer("o", "ø", "c", "k", "th", "t", "w", "z")
	return r.Replace(s)
}
