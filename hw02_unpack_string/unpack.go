package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(s string) (string, error) {
	if s == "" {
		return "", nil
	}

	sl := strings.Split(s, "")
	n := len(sl)

	if n > 1 && isNumeric(sl[0]) {
		return "", ErrInvalidString
	}

	sb := strings.Builder{}
	for i := 0; i < n; i++ {
		v := sl[i]
		if isBackslash(v) {
			if isLetter(getNext(sl, i)) {
				return "", ErrInvalidString
			}

			sb.WriteString(getNext(sl, i))
			i++
			continue
		}

		if isNumeric(v) {
			if isNumeric(getPrev(sl, i)) && isNumeric(getPrev(sl, i-1)) {
				return "", ErrInvalidString
			}

			cnt := getInt(sl, i)
			if cnt == 0 {
				return "", ErrInvalidString
			}

			sb.WriteString(strings.Repeat(getPrev(sl, i), cnt-1))
			continue
		}

		if isZero(getNext(sl, i)) {
			i++
			continue
		}

		sb.WriteString(v)
	}

	return sb.String(), nil
}

func isNumeric(s string) bool {
	return s >= "0" && s <= "9"
}

func isZero(s string) bool {
	return s == "0"
}

func isLetter(s string) bool {
	return !isNumeric(s) && !isBackslash(s)
}

func isBackslash(s string) bool {
	return s == "\\"
}

func getPrev(sl []string, from int) string {
	if from > 0 {
		return sl[from-1]
	}

	return ""
}

func getNext(sl []string, from int) string {
	if from < len(sl)-1 {
		return sl[from+1]
	}

	return ""
}

func getInt(sl []string, idx int) int {
	cnt, _ := strconv.Atoi(sl[idx])
	return cnt
}
