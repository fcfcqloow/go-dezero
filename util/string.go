package util

func Ellipsis(text string, length int) string {
	if r := []rune(text); len(r) > length {
		return string(r[0:length]) + "…"
	}
	return text
}
