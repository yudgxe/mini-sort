package cstrings

import "strings"

func Split(s string, sep string) []string {
	buff := make([]string, 0)
	i := 0

	strings := strings.Split(s, sep)
	for _, v := range strings {
		if v == "" {
			i++
		} else {
			buff = append(buff, AddRune(v, ' ', i))
			i = 0
		}
	}

	return buff
}

func AddRune(s string, r rune, size int) string {
	var sb strings.Builder

	for i := 0; i < size; i++ {
		sb.WriteRune(r)
	}

	sb.WriteString(s)
	return sb.String()
}

func BuildFromWords(sep string, args ...string) string {
	var sb strings.Builder
	size := len(args)-1

	for i := 0; i < size; i++ {
		sb.WriteString(args[i] + sep)
	}

	sb.WriteString(args[size])

	return sb.String()
}
