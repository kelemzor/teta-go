package main

import (
	"errors"
	"fmt"
	"regexp"
	"sort"
	"strconv"
)

var IsLetter = regexp.MustCompile(`^[a-zA-Zа-яА-Я]`).MatchString

func convolve_s(s *string) (string, error) {
	if !IsLetter(*s) {
		return "", errors.New("invalid string")
	}

	m := make(map[rune]int)

	for _, c := range *s {
		m[c]++
	}

	keys := make([]int, 0, len(m))
	for k := range m {
		keys = append(keys, int(k))
	}
	sort.Ints(keys)

	res := ""
	for _, k := range keys {
		c := rune(k)
		res += string(c) + strconv.Itoa(m[c])
	}

	return res, nil
}

func main() {
	in_s := ""

	fmt.Println("enter letters")
	fmt.Scan(&in_s)

	res, err := convolve_s(&in_s)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(res)
}
