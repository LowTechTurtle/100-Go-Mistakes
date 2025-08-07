package main

import "strings"

func concat1(s []string) string {
	var sCat string
	for _, v := range s {
		sCat += v
	}
	return sCat
}

func concat2(s []string) string {
	sb := strings.Builder{}
	for _, v := range s {
		sb.WriteString(v)
	}
	return sb.String()
}

func concat3(s []string) string {
	sb := strings.Builder{}
	var n int
	for _, v := range s {
		n += len(v)
	}
	sb.Grow(n)
	for _, v := range s {
		sb.WriteString(v)
	}
	return sb.String()
}
