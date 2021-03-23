package converter

import (
	"regexp"
	"strconv"
	"strings"
)

func StringToInt64(s string) (i int64) {
	re := regexp.MustCompile("[0-9]+")
	i, _ = strconv.ParseInt(re.FindString(s), 10, 64)
	return
}

func StringToFloat64(s string) (f float64) {
	re := regexp.MustCompile(`[0-9]+(\.[0-9]+)?`)
	f, _ = strconv.ParseFloat(re.FindString(s), 64)
	return
}

func StringToIntSlice(s string) []int64 {
	var list []int64
	for _, v := range strings.Split(strings.TrimSpace(s), ",") {
		i, _ := strconv.ParseInt(v, 10, 64)
		if i != 0 {
			list = append(list, i)
		}
	}
	return list
}

func StringToStringSlice(s string) []string {
	var arr []string
	for _, v := range strings.Split(strings.TrimSpace(s), ",") {
		if v != "" {
			arr = append(arr, v)
		}
	}
	return arr
}

func UniqueString(s []string) []string {
	m := make(map[string]struct{}, len(s))
	i := 0
	for _, t := range s {
		if _, ok := m[t]; ok {
			continue
		}
		m[t] = struct{}{}
		s[i] = t
		i++
	}
	return s[:i]
}

func UniqueInt64(s []int64) []int64 {
	m := make(map[int64]struct{}, len(s))
	i := 0
	for _, t := range s {
		if _, ok := m[t]; ok {
			continue
		}
		m[t] = struct{}{}
		s[i] = t
		i++
	}
	return s[:i]
}

func Contains(fields []string, field string) bool {
	for _, f := range fields {
		if f == field {
			return true
		}
	}
	return false
}
