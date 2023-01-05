package exce

import "sort"

func Sorted(s []string) []string {
	sort.Strings(s)
	return s
}